package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Hamza")
	// Unused variables aren't allowed in Go

	var name string = "Hamza"
	fmt.Printf("This is my name: %s\n", name)

	age := 27
	fmt.Printf("This is my age: %d\n", age)

	var city string
	city = "Amman"
	fmt.Printf("This is my city: %s\n", city)

	var country, continent string = "Jordan", "Asia"
	fmt.Printf("%s is based in %s which is in %s\n", city, country, continent)

	// To declare different variables with different types
	var (
		isEmployed bool = true
		salary int = 2500
		position string = "DevOps"
	)
	fmt.Printf("%s employed: %t, his salary equal %d and his position is %s\n", name, isEmployed, salary, position)

	// Zero values
	var defaultInt int
	fmt.Println(defaultInt)
}