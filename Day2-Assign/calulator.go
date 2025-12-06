package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	if operator == "+"{
		fmt.Print("Answer : ", a+b)
	} else if operator == "-"{
		fmt.Print("Answer : ", a-b)
	} else if operator == "*"{
		fmt.Print("Answer : ", a*b)
	} else if operator == "/"{
		if b != 0 {
			fmt.Print("Answer : ", a/b)
		} else {
			fmt.Print("Error: Division by zero")
		}
	} else {
		fmt.Print("Invalid operator")
	}
}

	