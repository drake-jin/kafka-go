package producer

import (
	"encoding/json"
	"flag"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
)

var (
	addr    = flag.String("addr", ":8080", "The address to bind to")
	brokers = flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
)

type user struct {
	Username string `json:"Username"`
	Role     string `json:"Role"`

	encoded []byte
	err     error
}

func (u *user) ensureEncoded() {
	if u.encoded == nil && u.err == nil {
		u.encoded, u.err = json.Marshal(u)
	}
}

func (u *user) Length() int {
	u.ensureEncoded()
	return len(u.encoded)
}

func (u *user) Encode() ([]byte, error) {
	u.ensureEncoded()
	return u.encoded, u.err
}

func Start() {
	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	producer := GetProducer(brokerList)
	entry := &user{
		Username: "DrakeJin",
		Role:     "Admin",
	}
	producer.SendMessage(&sarama.ProducerMessage{
		Topic: "important",
		Value: entry,
	})
	producer.Close()
}

func GetProducer(brokerList []string) sarama.SyncProducer {
	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	return producer
}
