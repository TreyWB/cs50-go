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

	fmt.Print("Input: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var vowels = map[string]string{
		"a": "",
		"e": "",
		"i": "",
		"o": "",
		"u": "",
	}

	var result = input
	for vowel, replacement := range vowels {
		result = strings.ReplaceAll(result, vowel, replacement)
	}

	fmt.Println("Output:", result)
}
