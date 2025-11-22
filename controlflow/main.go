package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("controlflow in go")

	fmt.Println("for loop in go example")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Day %d is %s\n", i, day)
	// }

	// for index, days := range days {
	// 	fmt.Printf("Day %d is %s\n", index, days)
	// }

	for _, day := range days {
		fmt.Println("Day is:", day)
	}

	// // break contine goto example
	// rougeValue := 5
	// for rougeValue < 10 {
	// 	fmt.Println("Rouge value is:", rougeValue)
	// 	rougeValue++

	// 	// if rougeValue == 7 {
	// 	// 	break
	// 	// }

	// 	// if rougeValue == 7 {
	// 	// 	rougeValue++
	// 	// 	continue
	// 	// }

	// 	if rougeValue == 7 {
	// 		goto skip

	// 	}

	// }

	// skip:
	// 	fmt.Println("Skipped printing rouge value 7")

	// // switch cas in go example
	// // dice game in go

	// rand.Seed(time.Now().UnixNano())

	// diceNumber := rand.Intn(6) + 1 // generates a number between 1 and 6

	// fmt.Println("Dice number rolled:", diceNumber)
	// switch diceNumber {
	// case 1:
	// 	fmt.Println("You rolled a One! Move 1 step forward.")
	// case 2:
	// 	fmt.Println("You rolled a Two! Move 2 steps forward.")
	// case 3:
	// 	fmt.Println("You rolled a Three! Move 3 steps forward.")
	// case 4:
	// 	fmt.Println("You rolled a Four! Move 4 steps forward.")
	// case 5:
	// 	fmt.Println("You rolled a Five! Move 5 steps forward.")
	// case 6:
	// 	fmt.Println("You rolled a Six! Move 6 steps forward and roll again!")
	// default:
	// 	fmt.Println("Invalid dice number!")
	// }

	// // if and else example
	// loginAccount := 20
	// if loginAccount < 10 {
	// 	fmt.Println("Login account is less than 10")
	// } else if loginAccount < 20 {
	// 	fmt.Println("Login account is less than 20")
	// } else {
	// 	fmt.Println("Login account is greater than or equal to 20")
	// }

	// if num := 9; num%2 == 0 {   // initialization statement
	// 	fmt.Println(num, "is even")
	// } else {
	// 	fmt.Println(num, "is odd")
	// }
}
