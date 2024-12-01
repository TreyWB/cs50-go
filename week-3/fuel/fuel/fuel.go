package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errInvalidInput = errors.New("error: invalid input")

func main() {
	var remainingFuel float64
	var err error

	for {
		remainingFuel, err = parseFraction()
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	switch {
	case remainingFuel <= 1:
		fmt.Println("E")
	case remainingFuel >= 99:
		fmt.Println("F")
	default:
		output := strconv.Itoa(int(remainingFuel)) + "% remaining"
		fmt.Printf("\n%v", output)
	}
}

func parseFraction() (float64, error) {
	var input string

	fmt.Println("Fraction:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := strings.SplitN(input, "/", 2)
	if len(result) != 2 {
		return 0.0, errInvalidInput
	}

	var x, y float64
	var err error

	if x, err = strconv.ParseFloat(result[0], 64); err != nil {
		return 0.0, errInvalidInput
	}
	if y, err = strconv.ParseFloat(result[1], 64); err != nil {
		return 0.0, errInvalidInput
	}

	if y == 0 {
		return 0.0, fmt.Errorf("error: denominator cannot be zero")
	} else if x > y {
		return 0.0, fmt.Errorf("error: numerator cannot be greater than denominator")
	}

	return x / y * 100, nil
}
