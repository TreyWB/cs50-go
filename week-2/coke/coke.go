package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var price = 50

	for price > 0 {
		fmt.Println("\nAmount Due:", price)

		coin, err := validateCoin()
		if err != nil {
			fmt.Println(err)
			continue
		}

		price = price - coin
	}

	change := price / -1
	fmt.Println("\nChange Due:", change)
}

func validateCoin() (int, error) {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insert a coin: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var coins = map[string]bool{
		"1":  true,
		"5":  true,
		"10": true,
		"25": true,
		"50": true,
	}

	if !coins[input] {
		return 0, errors.New("\ninvalid coin")
	}

	coinValue, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("\n error getting coin value")
	}

	return coinValue, nil
}
