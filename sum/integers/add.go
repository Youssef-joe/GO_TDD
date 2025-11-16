package integers

import "fmt"

// add two integers and return the sum of them
func Add(x, y int) int {
	return x + y
}

// multiply two integers and return the result
func Multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(Add(2, 4))
}
