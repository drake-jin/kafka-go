package producer

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println("producer start")
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
