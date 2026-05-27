package main

import "fmt"

// Struct definition
type Student struct {
	name string
	age  int
}

func main() {

	// Struct usage
	s1 := Student{
		name: "Kiruthika",
		age:  21,
	}

	fmt.Println("Name:", s1.name)
	fmt.Println("Age:", s1.age)
}
