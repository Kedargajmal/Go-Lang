package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	numbers = append(numbers, 6, 7, 8)
	fmt.Println("After appending elements:", numbers)
	fmt.Printf("Number has data tyoe : %T\n", numbers)

	fmt.Println("-----------------------------------")

	// Creating a slice using make()
	num := make([]int, 3, 5)
	fmt.Println("New slice with make():", num)
	fmt.Println("Length of num:", len(num))
	fmt.Println("Capacity of num:", cap(num))

	fmt.Println("-----------------------------------")

	// Appending elements to the slice created with make()
	num = append(num, 4,5,6)
	fmt.Println("After appending elements to num:", num)
	fmt.Println("Length of num after appending:", len(num))
	fmt.Println("Capacity of num after appending:", cap(num))	
}
