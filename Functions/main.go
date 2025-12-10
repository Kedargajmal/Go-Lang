package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("This is Main Functin")

	add := Add(4, 6)
	fmt.Println(add)
}
