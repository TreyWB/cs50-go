package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var names []string
	for {
		input := getInput()
		if input == "quit" || input == "exit" || input == "" {
			break
		}

		names = append(names, input)
	}

	builder := strings.Builder{}
	builder.WriteString("Adieu, adieu, to ")

	switch len(names) {
	case 1:
		builder.WriteString(names[0])
	case 2:
		builder.WriteString(names[0])
		builder.WriteString(" and ")
		builder.WriteString(names[1])
	default:
		for _, name := range names[:len(names)-1] {
			builder.WriteString(name)
			builder.WriteString(", ")
		}
		builder.WriteString("and ")
		builder.WriteString(names[len(names)-1])
	}

	output := builder.String()
	fmt.Println("Output:", output)
}

func getInput() string {
	var input string
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
