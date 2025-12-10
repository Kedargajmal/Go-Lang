package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Hello, World!"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two one"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str1 := "  hello, Go   "
	trimmed := strings.TrimSpace(str1)
	fmt.Println(trimmed)

	
}
