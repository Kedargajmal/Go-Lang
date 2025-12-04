//take input from user and print it
package main
import "fmt"

func Test() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("Hello", name, "your age is", age)
}

func main() {
	Test()
}