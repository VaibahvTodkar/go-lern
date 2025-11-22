package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in go-lang!")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages map:", languages)
		fmt.Println("Languages map:", languages)
	fmt.Println("Languages map:", languages)
	fmt.Println("Languages map:", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Languages map after deletion:", languages)

	// loops are interesting in go

	for key, value := range languages {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

// 	for _, value := range languages { // ignore key i dont know the data type of the key  or called error handling using _ symbol
// 		fmt.Printf("Key: %s Value: %s\n", value)
// 	}
}
