package main
import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := "Hello, Go!"
	for index, char := range data {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}