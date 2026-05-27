package main

import (
	"fmt"
	"sync"
	"time"
)

// Sleep sort function
func sleepSort(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Sleep based on number value
	time.Sleep(time.Duration(num) * time.Second)

	fmt.Println(num)
}

func main() {

	numbers := []int{3, 1, 4, 2}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)

		// Run goroutine
		go sleepSort(num, &wg)
	}

	wg.Wait()

	fmt.Println("Sleep sort completed")
}
