package services

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RemoveFunction(cmd *cobra.Command, args []string) {
	err := os.Remove("README.md")
	if err != nil {
		fmt.Println("Error removing README.md:", err)
		return
	}
	fmt.Println("README.md removed successfully!")
}
