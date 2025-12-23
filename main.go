package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "hello",
		Short: "A simple hello world command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Readmify!")
		},
	}

	rootCmd.Execute()
}
