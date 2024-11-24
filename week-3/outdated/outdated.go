package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errInvalidDate = fmt.Errorf("error: invalid date")

func main() {
	var day, month, year int
	var err error
	for {
		input := getDate()

		if strings.Contains(input, "/") {
			day, month, year, err = parseShortDate(input)
		} else if strings.Contains(input, ",") {
			day, month, year, err = parseLongDate(input)
		}

		if err != nil {
			continue
		}

		if !checkDate(day, month, year) {
			continue
		}

		break
	}

	output := fmt.Sprintf("%04d-%02d-%02d\n", year, month, day)
	fmt.Println(output)
}

func parseLongDate(input string) (int, int, int, error) {
	var parsed []string
	var months = map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}

	if strings.Contains(input, ",") {
		parsed = strings.SplitN(input, " ", 3)
	} else {
		return 0, 0, 0, errInvalidDate
	}

	if len(parsed) != 3 {
		return 0, 0, 0, errInvalidDate
	}

	var day, month, year int
	var err error

	if value, ok := months[parsed[0]]; ok {
		month = value
	} else {
		return 0, 0, 0, errInvalidDate
	}

	parsed_day := strings.Replace(parsed[1], ",", "", -1)
	if day, err = strconv.Atoi(parsed_day); err != nil {
		return 0, 0, 0, errInvalidDate
	}
	if year, err = strconv.Atoi(parsed[2]); err != nil {
		return 0, 0, 0, errInvalidDate
	}

	if !checkDate(day, month, year) {
		return 0, 0, 0, errInvalidDate
	} else {
		return day, month, year, nil
	}
}

func parseShortDate(input string) (int, int, int, error) {
	var parsed []string

	if strings.Contains(input, "/") {
		parsed = strings.SplitN(input, "/", 3)
	} else {
		return 0, 0, 0, errInvalidDate
	}

	if len(parsed) != 3 {
		return 0, 0, 0, errInvalidDate
	}

	var day, month, year int
	var err error

	if month, err = strconv.Atoi(parsed[0]); err != nil {
		return 0, 0, 0, errInvalidDate
	}
	if day, err = strconv.Atoi(parsed[1]); err != nil {
		return 0, 0, 0, errInvalidDate
	}
	if year, err = strconv.Atoi(parsed[2]); err != nil {
		return 0, 0, 0, errInvalidDate
	}

	if !checkDate(day, month, year) {
		return 0, 0, 0, errInvalidDate
	} else {
		return day, month, year, nil
	}
}

func checkDate(day int, month int, year int) bool {
	if day <= 0 || month <= 0 || year <= 0 {
		return false
	}

	if day > 31 || month > 12 {
		return false
	}

	return true
}

func getDate() string {
	var input string

	fmt.Println("\nEnter a date:")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input
}
