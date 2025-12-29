package services

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func InitFunction(cmd *cobra.Command, args []string) {
	structure, _ := cmd.Flags().GetBool("structure")
	name, _ := cmd.Flags().GetString("name")
	debug, _ := cmd.Flags().GetBool("debug")
	if debug {
		fmt.Println("Debug Info:")
		fmt.Print(string(getStructure()))
		return
	}

	// Fail if README exists
	destFile, err := os.OpenFile("README.md", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("README.md already exists. Initialization aborted.")
			return
		}
		fmt.Println("Error creating README:", err)
		return
	}
	defer destFile.Close()

	content := InitTemplate
	content = appendName(content, name)

	if structure {
		content = appendStructure(content, name)
		fmt.Println("Project Structure will be added to README")
	}

	// Write README
	_, err = destFile.Write(content)
	if err != nil {
		fmt.Println("Error writing README:", err)
		return
	}

	fmt.Println("README.md created successfully!")
}

func appendName(content []byte, name string) []byte {
	if name != "" {
		content = bytes.ReplaceAll(content, []byte("{{PROJECT_NAME}}"), []byte(name))
	} else {
		content = bytes.ReplaceAll(content, []byte("{{PROJECT_NAME}}"), []byte("Project Title"))
	}
	return content
}

func appendStructure(content []byte, name string) []byte {
	content = append(content, []byte("\n\n## Project Structure\n\n```text\n")...)
	content = append(content, getStructure()...)
	content = append(content, []byte("```\n")...)
	return content
}

func getStructure() []byte {
	var buf bytes.Buffer

	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip hidden
		if strings.Split(path, string(os.PathSeparator))[0] == ".git" ||
			strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		depth := len(strings.Split(path, string(os.PathSeparator))) - 1
		name := filepath.Base(path)

		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if depth == 0 {
			buf.WriteString("├── " + name + "\n")
		} else {
			prefix := strings.Repeat("│   ", depth) + "├── "
			buf.WriteString(prefix + name + "\n")
		}
		return nil
	}

	_ = filepath.Walk("./", visit)
	buf.WriteString("└── README.md\n")

	return buf.Bytes()
}
