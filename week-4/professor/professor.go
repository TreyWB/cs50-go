package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var level int
	var guess int
	var score int

	key := map[int]int{
		1: 10,
		2: 100,
		3: 1000,
	}

	for level <= 0 || level > 3 {
		level = getLevel()
	}

	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		x := r.Intn(key[level])
		y := r.Intn(key[level])

		for attempt := 0; attempt <= 4; attempt++ {
			if attempt == 3 {
				fmt.Printf("%v + %v = %v\n", x, y, x+y)
				break
			}

			fmt.Printf("%v + %v = ", x, y)
			guess = getGuess()

			if guess != x+y {
				fmt.Println("EEE")
			} else {
				score += 1
				break
			}
		}
	}

	fmt.Println("Score:", score)
}

func getLevel() int {
	var level int
	fmt.Print("Level: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	level, _ = strconv.Atoi(input)

	return level
}

func getGuess() int {
	var guess int
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	guess, _ = strconv.Atoi(input)

	return guess
}
