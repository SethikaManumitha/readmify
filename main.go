package main

import (
	"fmt"
	"readmify/services"

	"github.com/spf13/cobra"
)

func main() {
	initCmd := cobra.Command{
		Use:   "init",
		Short: "Initialize a new README file",
		Long:  "Initalize an empty README file",
		Run:   services.InitFunction,
	}

	initCmd.Flags().BoolP("structure","s",false,"Add Project Structure to README")
	rootCmd := cobra.Command{
		Use:   "readmify",
		Short: "A tool to generate README files",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Readmify!")
		},
	}

	rootCmd.AddCommand(&initCmd)
	rootCmd.Execute()
}
