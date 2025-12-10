package main
import "fmt"

func main() {

	// Declare and initialize an array of integers
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	//Accessing array elements using index
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// Modifying an array element
	numbers[1] = 25
	fmt.Println("Modified second element:", numbers[1])

	// Iterating over the array using a for loop
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Iterating over the array using range
	fmt.Println("Array elements using range:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Getting the length of the array
	fmt.Println("Length of the array:", len(numbers))

	// Multi-dimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi-dimensional array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Array initialization with shorthand syntax
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits array:", fruits)

	// Partial array initialization
	var partialArray [5]int = [5]int{1, 2}
	fmt.Println("Partially initialized array:", partialArray)

	// Array with inferred size
	inferredArray := [...]float64{3.14, 1.59, 2.65, 5.35}
	fmt.Println("Inferred size array:", inferredArray)

	// Copying arrays
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 10
	fmt.Println("Original array after copying and modifying copied array:", original)
	fmt.Println("Copied array:", copied)

	//Array comparison
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	array3 := [3]int{4, 5, 6}
	fmt.Println("array1 == array2:", array1 == array2)
	fmt.Println("array1 == array3:", array1 == array3)
	

}
