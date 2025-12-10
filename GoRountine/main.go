package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("1")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2")
}

func sayHi() {
	fmt.Println("3")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("4")
}

func main() {
	fmt.Println("Go Rountine")
	go sayHello()
	go sayHi()

	//wait for moment to allow the goroutine to finish
	time.Sleep(3000 * time.Millisecond)
}
