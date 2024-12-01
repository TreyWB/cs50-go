package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dustin/go-humanize/english"
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
	builder.WriteString(english.OxfordWordSeries(names, "and"))
	fmt.Println(builder.String())
}

func getInput() string {
	var input string
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
