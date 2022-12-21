package main

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	fmt.Println("Hello")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Println("Error", err)
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// fmt.Println(numRating + 1)
}
