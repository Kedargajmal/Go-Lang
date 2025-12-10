package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time is :", currentTime)

	formatted := currentTime.Format("02-01-2006")
	fmt.Println("Formatted time is : ", formatted)

	layoutstr := "2006-01-02"
	datastr := "2025-12-10"
	formattedTime, _ := time.Parse(layoutstr, datastr)
	fmt.Println(formattedTime)
}
