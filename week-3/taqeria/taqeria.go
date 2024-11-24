package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var errInvalidItem = errors.New("error: invalid menu item")

func main() {
	var total float64

	var menu = map[string]float64{
		"Baja Taco":        4.25,
		"Burrito":          7.50,
		"Bowl":             8.50,
		"Nachos":           11.00,
		"Quesadilla":       8.50,
		"Super Burrito":    8.50,
		"Super Quesadilla": 9.50,
		"Taco":             3.00,
		"Tortilla Salad":   8.00,
		"Done":             0.00,
		"Exit":             0.00,
		"Quit":             0.00,
	}
	for item, value := range menu {
		if item != "Done" && item != "Exit" && item != "Quit" {
			fmt.Printf("%v: $%.2f\n", item, value)
		}
	}

	for {
		input := getItem()

		if value, ok := menu[input]; ok {
			if value == 0.00 {
				break
			}

			total += value
			fmt.Printf("\nTotal: $%.2f", total)
		} else {
			fmt.Println(errInvalidItem)
		}
	}
}

func getItem() string {
	var input string

	fmt.Println("\nSelect Item:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = cases.Title(language.AmericanEnglish, cases.NoLower).String(input)

	return input
}
