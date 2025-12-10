// take input from user and print it
package main

import (
	"bufio"
	"fmt"
	"os"
)

func Test() {
	// var name string
	// var age int

	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)

	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&age)

	// fmt.Println("Hello", name, "your age is", age)

	//USING BUFIO PACKAGE
	fmt.Println("What is your name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)

}
