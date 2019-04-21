package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type scaler interface {
	scale(f float64)
}

type Vertex struct {
	X, Y float64
}

type FloatValue float64

/*
	Like other OOP, in GO an interface type variable can hold the value of recevier which
	has implemented the interface.
	If the interface is nil, the method will be called with nil receiver

*/

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f FloatValue) Abs() float64 {
	if v := float64(-f); f < 0 {
		return v
	}
	return float64(f)
}

func (v *Vertex) scale(f float64) {

}

func (fv *FloatValue) scale(f float64) {

}

/*
Empty interface: Empty interface contains no method.
*/
func emptyInterface() {
	var i interface{}
	i = 10
	describe(i)
	i = "shihab"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("Type : %T, Value: %v", i, i)
}

func main() {
	var a Abser
	v := Vertex{1, 2}
	f := FloatValue(-2)
	a = f
	fmt.Println("Calling to the FloatValue reciver method: ", a.Abs())
	a = &v
	// a = v will give error as Vertex doesn't implement Abs(). *Vertex implement Abs()
	fmt.Println("Calling to the Vertex receiver method : ", a.Abs())

	fmt.Println("***************** Empty Interface *************")
	emptyInterface()
	fmt.Println("***************** Type Assertion *************")
	typeAssert("shihab")
}
