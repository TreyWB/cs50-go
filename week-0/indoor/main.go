package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var original string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a sentence:")
	original, _ = reader.ReadString('\n')
	original = strings.TrimSpace(original)
	original = strings.ToLower(original)

	fmt.Println("\n" + original)
}
