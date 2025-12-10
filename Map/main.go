package main
import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Alice"] = 85
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 78
	studentGrades["Diana"] = 92
	fmt.Println("Mark of Bob : ", studentGrades["Bob"])

	studentGrades["Alice"] = 95
	fmt.Println("Updated mark of Alice : ", studentGrades["Alice"])

	delete(studentGrades, "Alice")
	fmt.Println("Updated student grades: ", studentGrades)

	//Checking if a key exists
	grade, exists := studentGrades["Charlie"]
	if exists {
		fmt.Println("Charlie's grade is:", grade)
	} else {
		fmt.Println("Charlie not found in the map.")
	}

	for index, value := range studentGrades {
		fmt.Printf("Student: %s, Grade: %d\n", index, value)
	}
}