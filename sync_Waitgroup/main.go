package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d started\n", i)
	//some task is happening
	fmt.Printf("worker %d end", i)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	//wait for all worker to finish
	wg.Wait()

	fmt.Println("worker task completed")
}
