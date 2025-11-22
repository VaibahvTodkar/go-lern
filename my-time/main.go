package main

import (
	"fmt" 
	"time"
)

func main() {
	fmt.Println("Hello, World in Time!")

	presentTime := time.Now()
	fmt.Println("Current Time is:", presentTime)

	fmt.Println("Formatted Time:", presentTime.Format("2006-01-02 15:04:05"))

	createdDate := time.Date(2020, time.January, 1, 10, 0, 0, 0, time.UTC)
	fmt.Println("Created Date:", createdDate)

	fmt.Println("Created Date Formatted:", createdDate.Format("2006-01-02 15:04:05 Monday "))

}

