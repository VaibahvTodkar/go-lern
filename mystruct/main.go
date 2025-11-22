package main

import (
	"fmt"
)

func main() {
	fmt.Println("My Structs in go-lang!")
	// no inhertance in go-lang
	// no class in go-lang
	// no super parent in go-lang

	// creating struct object
	user1 := User{
		Name:   "John Doe",
		Email:  "john.deo@gmail.com",
		Status: true,
		Age:    30,
	}
	fmt.Println("User 1:", user1)
	fmt.Printf("User 1 Details: %+v\n", user1)
	fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", user1.Name, user1.Email, user1.Status, user1.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
