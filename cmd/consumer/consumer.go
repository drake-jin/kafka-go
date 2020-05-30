package consumer

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println("consumer start")
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
