package cmd

import (
	"github.com/drake-jin/kafka-go/cmd/consumer"
	"github.com/drake-jin/kafka-go/cmd/producer"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := cobra.Command{
		Use: "kafka-go",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(producer.GetCommand())
	cmd.AddCommand(consumer.GetCommand())
	cmd.Execute()
}
