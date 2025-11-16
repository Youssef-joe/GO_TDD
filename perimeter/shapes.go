package perimeter

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// here that's not a function actually it's a method
// the key change here is that methods acts as a receiver
// The only difference is the syntax of the method receiver
// func (receiverName ReceiverType) MethodName(args).
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}