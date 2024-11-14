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

	fmt.Println("What time is it?")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	time, err := convert(input)
	if err != nil {
		fmt.Println("\nError:", err)
		return
	}

	if time >= 7.0 && time <= 8.0 {
		fmt.Println("\nBreakfast Time!")
	} else if time >= 12.0 && time <= 13.0 {
		fmt.Println("\nLunch Time!")
	} else if time >= 18.0 && time <= 19.0 {
		fmt.Println("\nDinner Time!")
	} else {
		fmt.Println("\nFAT!")
	}
}

func convert(time string) (float64, error) {
	result := strings.Split(time, ":")
	if len(result) != 2 {
		return 0, fmt.Errorf("invalid time")
	}

	var hours float64
	var minutes float64
	var err error

	if hours, err = strconv.ParseFloat(result[0], 64); err != nil {
		return 0, fmt.Errorf("invalid time")
	}
	if minutes, err = strconv.ParseFloat(result[1], 64); err != nil {
		return 0, fmt.Errorf("invalid time")
	}

	return hours + minutes/60, nil
}
