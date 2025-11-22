package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, Slices!")

	var fruitList = []string{"Apple", "Banana", "Grapes"}
	fmt.Println("Fruit List:", fruitList)
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("Updated Fruit List:", fruitList)

	fruitList = fruitList[1:3]
	fmt.Println("Sliced Fruit List:", fruitList)

	highScore := make([]int, 4) // length 4
	highScore[0] = 234
	highScore[1] = 456
	highScore[2] = 678
	highScore[3] = 890
	// highScore[4] = 1000 // This will cause an error: index out of range

	fmt.Println("High Scores:", highScore)

	highScore = append(highScore, 1000, 1200, 1400)

	fmt.Println("Updated High Scores:", highScore)
	
	sort.Ints(highScore)
	fmt.Println("Sorted High Scores:", highScore)


	// remove value based on index 

	var cousineList = []string{"Aman", "Suman", "Raman", "Kaman"}
	fmt.Println("Cousine List:", cousineList)

	var indexToRemove int = 2
	cousineList = append(cousineList[:indexToRemove], cousineList[indexToRemove+1:]...)
	fmt.Println("Updated Cousine List:", cousineList)
}
