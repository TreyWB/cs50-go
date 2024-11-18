package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	checkArgs()

	for {
		plate := getPlate()

		valid, err := checkPlate(plate)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if valid {
			fmt.Println("Valid")
			break
		} else {
			fmt.Println("Invalid")
			continue
		}
	}
}

func getPlate() string {
	var input string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nInput: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

func checkPlate(plate string) (bool, error) {
	match, err := regexp.MatchString("^[A-Z]{2,6}$|^[A-Z]{2,5}[1-9][0-9]?$", plate)
	if err != nil {
		return false, err
	}

	return match, nil
}

func checkArgs() {
	content := `
Use this prograom to check if the proposed vanity plate is valid.

Run the program without any arguments to begin.

Vanity Plate Requirements:
- Length: 2 to 6 characters
- Must start with at least two letters
- Number must be at the end of the plate
- No periods, spaces, or punctuation marks
`

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "--h" {
			fmt.Println(content)
			os.Exit(0)
		} else {
			fmt.Println("\nInvalid argument: Use -h, --h, or --help for help")
			os.Exit(0)
		}
	}
}
