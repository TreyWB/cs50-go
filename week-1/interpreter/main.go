package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nExpression:")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	x, y, z, err := getValues(input)
	if err != nil {
		fmt.Println("\nError:", err)
		return
	}

	result, err := evaluate(x, y, z)
	if err != nil {
		fmt.Println("\nError:", err)
		return
	}

	fmt.Printf("\n%0.1f", result)
}

func evaluate(x int, y string, z int) (float64, error) {
	switch y {
	case "+":
		return float64(x + z), nil
	case "-":
		return float64(x - z), nil
	case "*":
		return float64(x * z), nil
	case "/":
		return float64(x / z), nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func getValues(expression string) (x int, y string, z int, err error) {
	result := strings.Split(expression, " ")

	if len(result) != 3 {
		return 0, "", 0, fmt.Errorf("invalid expression")
	}

	if x, err = strconv.Atoi(result[0]); err != nil {
		return 0, "", 0, fmt.Errorf("x must be an integer")
	}
	if z, err = strconv.Atoi(result[2]); err != nil {
		return 0, "", 0, fmt.Errorf("z must be an integer")
	}

	operators := [4]string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if result[1] == operator {
			y = result[1]
			break
		}
	}

	if y == "" {
		return 0, "", 0, fmt.Errorf("invalid operator")
	}

	return x, y, z, err
}
