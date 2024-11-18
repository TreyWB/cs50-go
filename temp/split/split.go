package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "testLol"

	result := strings.SplitN(s, "L", 2)

	fmt.Println(result)
}
