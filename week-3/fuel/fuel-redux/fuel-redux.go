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
	var fraction *Fraction
	var err error

	for {
		fraction, err = getFraction()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err := fraction.Calculate(); err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	switch {
	case fraction.percent <= 1:
		fmt.Println("\nE")
	case fraction.percent >= 99:
		fmt.Println("\nF")
	default:
		remainingFuel := fmt.Sprintf("%v%% remaining", fraction.percent)
		fmt.Printf("\n%v", remainingFuel)
	}
}

type Fraction struct {
	numerator   float64
	denominator float64
	percent     float64
}

func getFraction() (*Fraction, error) {
	var input string

	fmt.Println("Fraction:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := strings.SplitN(input, "/", 2)
	if len(result) != 2 {
		return nil, errInvalidInput
	}

	var x, y float64
	var err error

	if x, err = strconv.ParseFloat(result[0], 64); err != nil {
		return nil, errInvalidInput
	}
	if y, err = strconv.ParseFloat(result[1], 64); err != nil {
		return nil, errInvalidInput
	}

	f := &Fraction{
		numerator:   x,
		denominator: y,
	}

	return f, nil
}

func (f *Fraction) Calculate() error {
	if f.denominator == 0 {
		return fmt.Errorf("error: denominator cannot be zero")
	}

	if f.numerator > f.denominator {
		return fmt.Errorf("error: numerator cannot be greater than denominator")
	}

	f.percent = f.numerator / f.denominator * 100
	return nil
}
