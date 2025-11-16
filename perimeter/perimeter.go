package perimeter

import "fmt"

func Perimeter(rec Rectangle) float64 {
	return 2*(rec.Width + rec.Height)
}
func Area(rec Rectangle)  float64 {
	return (rec.Width * rec.Height)
}

func perimeter () {
	fmt.Println("hello")
}