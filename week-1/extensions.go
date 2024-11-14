package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nFile Name:")
	input, _ = reader.ReadString('\n')
	file := strings.TrimSpace(input)

	extension, err := getFileExtension(file)
	if err != nil {
		fmt.Println("\nError:", err)
		return
	}

	var extension_list = map[string]string{
		"gif":  "image/gif",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"pdf":  "application/pdf",
		"txt":  "text/plain",
		"zip":  "application/zip",
	}

	var valid bool
	for key, value := range extension_list {
		if key == extension {
			fmt.Printf("\n%v", value)
			valid = true
			break
		}
	}

	if !valid {
		fmt.Println("\napplication/octet-stream")
	}
}

func getFileExtension(filename string) (string, error) {
	if result := strings.SplitAfter(filename, "."); len(result) > 1 {
		return result[1], nil
	} else {
		return "", fmt.Errorf("file contains no extension")
	}
}
