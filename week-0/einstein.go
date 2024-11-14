package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var original string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the value of m:")
	original, _ = reader.ReadString('\n')
	original = strings.TrimSpace(original)

	m, err := strconv.Atoi(original)

	if err != nil {
		fmt.Println("Error:", err)
	}

	e := formula(m)

	fmt.Printf("\n%v", e)
}

func formula(m int) int {
	var c int = 300000000

	e := float64(m) * math.Pow(float64(c), 2)

	return int(e)
}
