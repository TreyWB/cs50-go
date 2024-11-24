package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mpvl/unique"
)

func main() {
	var items []string
	groceries := make(map[string]int)
	fmt.Println("\nEnter Item List:")

	for {
		item := getItem()
		if item == "" {
			break
		}

		items = append(items, item)
	}

	for _, item := range items {
		if value, ok := groceries[item]; ok {
			groceries[item] = value + 1
		} else {
			groceries[item] = 1
		}
	}

	sort.Strings(items)
	unique.Strings(&items)
	for _, item := range items {
		fmt.Printf("%v %v\n", groceries[item], strings.ToUpper(item))
	}
}

func getItem() string {
	var input string

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input
}
