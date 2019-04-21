package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
	GO doesn't have class. But we can add receiver argument.
	The receiver argument between `func` keyword and method name.
	Adding recevier to a method means the method can be called from the receiver.
	In the value receiver example, From Vertex type v the Abs method can be called.
	In every method call the value is copied to the receiver end
*/

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// declare a type : type <typeName> <type>
type FloatValue float32

/*
	Example of non-struct receiver type with method.
	A method can not use a receiver which is defined in another package
*/

func (f FloatValue) Abs() float32 {
	if v := float32(-f); f < 0 {
		return v
	}
	return float32(f)
}

// pointer recevier.
/*
	Pointer recevier can be used in two cases:
	1. The method can modify the value that receiver points to
	2. Avoid copying the value each method call, useful when the struct contains may fields
*/

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// scale and scale1 method do the same thing
func scale1(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("Vertex abs value receiver: ", v.Abs())
	f := FloatValue(-1)
	fmt.Println("Float Abs: ", f.Abs())
	v.scale(.1) // Go interpret v.scale(.1) with (&v).scale(.1)
	fmt.Println("Vertex abs after pointer receiver: ", v.Abs())
}
