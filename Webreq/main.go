package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Time!")
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)

	//Read the body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Println("Response Body:", string(data))
	

}
