package main
import "fmt"

func main() {
	fmt.Println("Start of main function")
	defer fmt.Println("Deferred: This will run at the end of main function")
	fmt.Println("End of main function")
}