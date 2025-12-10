package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Printf("The integer value is: %d\n", num)
		fmt.Printf("Type of str is %T\n", num)


	var data float64 = float64(num)
	fmt.Printf("The float64 value is: %f\n", data)
		fmt.Printf("Type of str is %T\n", data)


	num = 123
	str := strconv.Itoa(num)
	fmt.Printf("The string value is: %s\n", str)
	fmt.Printf("Type of str is %T\n", str)

	number_String := "1234"
	number_int, _ := strconv.Atoi(number_String)
	fmt.Printf("The integer value is: %d\n", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

}