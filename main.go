package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()

	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("Letter: %c\n", ch)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()

	fmt.Println("Goroutines completed")
}
