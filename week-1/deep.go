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

	fmt.Println("What is the answer to Life, the Universe, and Everything?")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	switch input {
	case "42":
		fmt.Println("\nYes")
	case "forty-two":
		fmt.Println("\nYes")
	case "forty two":
		fmt.Println("\nYes")
	default:
		fmt.Println("\nNo")
	}
}
