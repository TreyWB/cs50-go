package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cost float64 = cost_input()
	var percentage float64 = percentage_input()

	tip := cost * percentage

	fmt.Printf("\nTip amount: $%.2f", tip)
}

func cost_input() float64 {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the cost of the meal:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.Contains(input, "$") {
		input = strings.ReplaceAll(input, "$", "")
	}

	cost, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return cost
}

func percentage_input() float64 {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the desired tip percentage:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.Contains(input, "%") {
		input = strings.ReplaceAll(input, "%", "")
	}

	percentage, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return percentage / 100
}
