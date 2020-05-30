package consumer

import (
	"github.com/spf13/cobra"

	saramaConsumer "github.com/drake-jin/kafka-go/internal/sarama/consumer"
)

func GetStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, _ []string) {
			saramaConsumer.Start()
		},
	}
	return cmd
}

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "consumer",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(GetStartCommand())
	return cmd
}
