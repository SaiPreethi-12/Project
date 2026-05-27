package main

import "fmt"

func displayMessage(name string) {
	fmt.Println("Hello", name)
}

func main() {

	// Variable
	age := 20

	// Condition
	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}

	// Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}

	// Function call
	displayMessage("Go User")
}
