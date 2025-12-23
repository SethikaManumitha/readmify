package services

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func InitFunction(cmd *cobra.Command, args []string) {
	structure, _ := cmd.Flags().GetBool("structure")

	// fails if file exists
	destFile, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {
		// Check if file already exists
		if os.IsExist(err) {
			fmt.Println("README.md already exists. Initialization aborted to prevent overwriting.")
			return
		} else {
			fmt.Println(err)
			return
		}
	}
	defer destFile.Close()

	// Check flag
	if structure {
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
