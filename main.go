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

	removeCmd := cobra.Command{
		Use:   "rm",
		Short: "Remove the existing README file",
		Long:  "Remove the existing README file from the current directory",
		Run:   services.RemoveFunction,
	}

	//initCmd.Flags().BoolP("debug", "d", false, "Temporary debug setting")

	initCmd.Flags().BoolP("structure", "s", false, "Add Project Structure to README")
	initCmd.Flags().StringP("name", "n", "", "Specify the project name")

	rootCmd := cobra.Command{
		Use:   "readmify",
		Short: "A tool to generate README files",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Readmify!")
		},
	}

	rootCmd.AddCommand(&initCmd)
	rootCmd.AddCommand(&removeCmd)
	rootCmd.Execute()
}
