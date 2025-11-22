package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Rating: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered Rating:", input)

	// Trim newline
	input = strings.TrimSpace(input)

	// Convert string to float
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error converting rating:", err)
		return
	}

	fmt.Println("Numeric Rating is:", numRating)
}
