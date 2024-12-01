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
	var answer int

	for level <= 0 {
		level = getLevel()
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer = r.Intn(level) + 1

	for guess != answer {
		guess = getGuess()

		if guess <= 0 {
			continue
		} else if guess > answer {
			fmt.Println("Too large!")
		} else if guess < answer {
			fmt.Println("Too small!")
		} else if guess == answer {
			fmt.Println("Just right!")
			break
		}
	}
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
	fmt.Print("Guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	guess, _ = strconv.Atoi(input)

	return guess
}
