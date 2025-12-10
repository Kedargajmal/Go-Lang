package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a new file
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()

	// // Write some content to the file
	// content := "Hello, World!"
	// _, errs := io.WriteString(file, content)
	// if errs != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }

	// fmt.Println("file created")



	// Open an existing file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the content of the file
	buffer := make([]byte, 1024)
	// Read until EOF
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Print the content to the console
		fmt.Print(string(buffer[:n]))
	}
}
