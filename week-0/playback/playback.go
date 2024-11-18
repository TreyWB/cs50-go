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
	original = strings.TrimSpace(original)

	modified = strings.Replace(original, " ", "...", -1)

	fmt.Println("\n" + modified)
}
