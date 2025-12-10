package main
import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0,
		fmt.Errorf("Denimonator can not be zero")
	}
	return a / b, nil
}

func main() {
	ans, _ := divide(10,0)
	fmt.Println(ans)
}