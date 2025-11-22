package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	imput, _ := reader.ReadString('\n')
	fmt.Println("You entered:", imput)
}
