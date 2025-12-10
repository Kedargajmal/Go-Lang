package main

import (
	"fmt"
)

func modifyValue(val *int) {
	*val = *val + 5
}

func main() {
	var num int = 42

	var ptr *int
	ptr = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value of ptr (address of num):", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr)

	// Modify the value using the pointer
	value := 10
	modifyValue(&value)
	fmt.Println("Modified value:", value)
}
