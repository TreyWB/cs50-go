package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/enescakir/emoji"
)

func main() {
	var input string
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	output := emoji.Parse(input)
	fmt.Println("Output:", output)
}
