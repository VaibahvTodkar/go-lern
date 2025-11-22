package main

import (
	"fmt"
)

func main() {
	fmt.Println("My Methos in go-lang!")
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
	user1.GetStatus()
	user1.GetAge()
	fmt.Println("Age after method call:", user1.Age) // Age will remain unchanged
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


func (u User) GetStatus()  {
	fmt.Println("is useractive", u.Status)
}

func (u User) GetAge() {
	u.Age = 10
	fmt.Println("Age in method:", u.Age)
}