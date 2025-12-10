package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
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
