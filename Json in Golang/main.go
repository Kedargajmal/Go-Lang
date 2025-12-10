package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func main() {
	User := Person{Name: "Alice", Age: 30, Email: "alice@123", Active: true}
	fmt.Println(User)

	//convert struct to json (marshal)
	jsonDate, err := json.Marshal(User)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User data is", string(jsonDate))

	//convert json to struct (unmarshal)
	jsonString := `{"name":"Bob","age":25,"email":"bob@123","active":false}`
	var User2 Person
	err = json.Unmarshal([]byte(jsonString), &User2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User2 data is", User2)

}
