package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var original string
	var modified string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a sentence:")
	original, _ = reader.ReadString('\n')

	modified = convert(strings.TrimSpace(original))

	fmt.Println("\n" + modified)
}

func convert(original string) string {
	replacements := map[string]string{
		":)": "🙂",
		":(": "🙁",
	}

	result := original
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}

	return result
}
