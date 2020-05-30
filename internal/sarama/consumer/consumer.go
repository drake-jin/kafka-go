package consumer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

// Sarama configuration options
var (
	brokers  = "127.0.0.1:9092"
	version  = "2.5.0"
	group    = "111example"
	topics   = "important"
	assignor = "range"
	oldest   = true
	verbose  = false
)

func Start() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	u := User{
		Username: "Hello",
		Role:     "User",
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(topics, ","), &u); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			//u.ready = make(chan bool)
		}
	}()

	//<-u.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

type User struct {
	Username string `json:"Username"`
	Role     string `json:"Role"`

	encoded []byte
	err     error
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (user *User) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("- Setup - about this record")
	// Mark the consumer as ready
	//close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (user *User) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("- Cleanup - about this record")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (user *User) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("- ConsumeClaim - about this record")
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, topic = %v",
			string(message.Value),
			message.Topic,
		)
		session.MarkMessage(message, "")
	}
	return nil
}
