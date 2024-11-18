package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("camelCase: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	result := camelToSnakeCase(input)
	fmt.Println("snake_case:", result)
}

func camelToSnakeCase(s string) string {
	var builder strings.Builder
	builder.Grow(len(s) + 4) // Preallocate some space for efficiency

	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			builder.WriteByte('_')
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
