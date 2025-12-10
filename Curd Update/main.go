package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	//convert json data to string
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

func performUpdateRequest() {
	todo := Todo{
		Userid:    222,
		Title:     "New 2nd Todo Item",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
		return
	}

	//convert json data to string
	jsonString := string(jsonData)
	fmt.Println("JSON Data:", jsonString)

	// convert json data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create put request
	req, err := http.NewRequest("PUT", myurl, jsonReader)
	if err != nil {
		panic(err)
		return
	}

	//set content-type header
	req.Header.Set("Content-Type", "application/json")

	//send the request using http client
	clent := http.Client{}
	res, err := clent.Do(req)
	if err != nil {
		panic(err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)
	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println("Response :", string(data))

}

func performDeleteRequest(){
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create put request
	req, err := http.NewRequest("DELETE", myurl, nil)
	if err != nil {
		panic(err)
		return
	}
	//set content-type header
	req.Header.Set("Content-Type", "application/json")

	//send the request using http client
	clent := http.Client{}
	res, err := clent.Do(req)
	if err != nil {
		panic(err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)
}

func main() {
	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()
}
