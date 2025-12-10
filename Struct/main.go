package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	Street string
	City   string
	state  string
}

type contact struct {
	Phone string
	Email string
}

type Employee struct {
	User_Details Person
	User_Address Address
	User_Contact contact
}

func main() {
	User1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println("User1:", User1)

	User2 := Person{"Jane", "Smith", 25}
	fmt.Println("User2:", User2)

	var employee1 Employee
	employee1.User_Details = Person{"Alice", "Johnson", 28}
	employee1.User_Address = Address{"123 Main St", "Springfield", "IL"}
	employee1.User_Contact = contact{"555-1234", "alice@1123"}

	fmt.Println("Employee1 Details:", employee1)

}
