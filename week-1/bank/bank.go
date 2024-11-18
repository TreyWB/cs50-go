package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	var result bool
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Greeting:")
	input, _ = reader.ReadString('\n')
	greeting := strings.TrimSpace(input)
	greeting = strings.ToLower(greeting)

	if _, result = strings.CutPrefix(greeting, "hello"); result {
		fmt.Println("$0")
	} else if _, result = strings.CutPrefix(greeting, "h"); result {
		fmt.Println("$20")
	} else {
		fmt.Println("$100")
	}
}
