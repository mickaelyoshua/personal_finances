package main

import "fmt"

func main() {
	errors := make(map[string]string)
	errors["name"] = "Name must be between 3 and 50 characters"
	fmt.Println("Initial map:", errors)
	fmt.Println("Length of map:", len(errors))
}