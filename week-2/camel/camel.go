package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("camelCase: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result, err := camelToSnakeCase(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("snake_case:", result)
}

func camelToSnakeCase(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input is empty")
	}
	if unicode.IsUpper(rune(input[0])) {
		return "", fmt.Errorf("input is not in camelCase")
	}

	var builder strings.Builder

	for i, r := range input {
		if unicode.IsUpper(r) && i > 0 {
			char := strings.ToLower(string(r))
			builder.WriteString("_")
			builder.WriteString(char)
		} else {
			builder.WriteString(string(r))
		}
	}

	return builder.String(), nil
}
