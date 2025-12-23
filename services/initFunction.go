package services

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func InitFunction(cmd *cobra.Command, args []string) {
	structre, _ := cmd.Flags().GetBool("structure")

	destFile, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Error creating README.md file:", err)
		return
	}
	defer destFile.Close()

	if structre {
		_, err = io.Copy(destFile, bytes.NewReader(StructureTemplate))
		fmt.Println("Project Structure will be added to README")
	} else {
		_, err = io.Copy(destFile, bytes.NewReader(InitTemplate))
	}

	if err != nil {
		fmt.Println("Error writing content:", err)
		return
	}

	fmt.Println("README.md created successfully!")
}
