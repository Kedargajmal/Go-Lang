package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	Userid    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-OK HTTP status")
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// 	return
	// }
	// fmt.Println("Response Body:", string(data))

	// Decode JSON response into Todo struct
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Todo Item: %+v\n", todo)
}

func performPostRequest() {
	todo := Todo{
		Userid:    23,
		Title:     "New Todo Item",
		Completed: false,
	}

	// Convert todo struct to JSON
	josnData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
		return
	}

	//convert jsin data to string
	jsonString := string(josnData)
	fmt.Println("JSON Data:", jsonString)

	// convert json data to reader
	jsonReader := strings.NewReader(jsonString)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		panic(err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)

	// Decode JSON response into Todo struct
	var createdTodo Todo
	err = json.NewDecoder(res.Body).Decode(&createdTodo)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Created Todo Item: %+v\n", createdTodo)

}

func main() {
	performGetRequest()
	performPostRequest()
}
