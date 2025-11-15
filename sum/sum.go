package sum

import "fmt"

// here that argument is called Variadic Functions which lets us introduce
// a variable ammount of arguments
func SumAll(numbersToSum ...[]int) []int {

	//we here can also create a slice with a capacity of "length"
	// sum := make([]int, length)
	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
		// but In this implementation, we are worrying less about capacity.
		//  We start with an empty slice sums and append to it
		//  the result of Sum as we work through the varargs.
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {

			//Slices can be sliced! The syntax is slice[low:high].
			// If you omit the value on one of the sides of the : it captures everything to that side of it.
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}

func Sum(numbers []int) int {
	res := 0

	for _, number := range numbers {
		res += number
	}
	return res
}

func main() {
	fmt.Println("Hello there")
}
