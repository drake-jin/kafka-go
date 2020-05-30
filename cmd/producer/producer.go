package producer

import (
	"github.com/spf13/cobra"

	saramaProducer "github.com/drake-jin/kafka-go/internal/sarama/producer"
)

/*
- 비동기 작성이 가능

*/
func GetStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, _ []string) {
			saramaProducer.Start()
		},
	}
	return cmd
}

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "producer",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(GetStartCommand())
	return cmd
}
