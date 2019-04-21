package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func doBasicOnPointer() {
	i, j := 10, 20

	fmt.Println("i: ", i)
	fmt.Println("j: ", j)

	p := &i
	fmt.Println("P's value: ", *p) // it'll show the value of i, as p refers to i's address
	*p = 15
	fmt.Println("i's new value: ", i) //it's called indirecting or dereferencing. Value is changed by a it's pointer
	p = &j
	*p = 30
	fmt.Println("j's new value : ", j) // simlar way, as p refer to j's address. indirection or dereferencing
	*p = *p / i
	fmt.Println("J's value after division: ", j) //
	fmt.Println("P's reference value : ", *p)
}

func getDistance(point1 Vertex, point2 Vertex) float64 {
	d := (point1.X-point2.X)*(point1.X-point2.X) + (point1.Y-point2.Y)*(point1.Y-point2.Y)
	return math.Sqrt(float64(d))
}

/*
	Logically we should write (*p).X to access X, but it looks cumbersome,
	p.X works.
*/
func strctWithPointer(v Vertex) {
	p := &v
	fmt.Println("Value of x : ", p.X)
	fmt.Println("Value of y : ", p.Y)
}

/*
 */

func structLiteralInit() {
	var (
		v1 = Vertex{1, 2}  // v1 struct with x=1 and y = 2
		v2 = Vertex{X: 1}  // v2 struct with x=1 and y = 0 (zero value)
		v3 = Vertex{}      // v3 struct with x=0 and y=0
		p  = &Vertex{1, 2} // p a pointer of struct with x=1 and y=2
	)

	fmt.Println("v1: ", v1, "v2:", v2, "v3:", v3, "p: ", p)
}
func main() {
	fmt.Println("Doing some basic operation using pointer: ")
	doBasicOnPointer()
	fmt.Println("Struct In action : ")
	fmt.Println(Vertex{1, 2})

	fmt.Println("Accessing struct element using dot: ")
	point1 := Vertex{2, 0}
	point2 := Vertex{0, 2}
	fmt.Println("Distance : ", getDistance(point1, point2))

	fmt.Println("Accessing struct using pointer: ")
	strctWithPointer(point1)
	fmt.Println("Struct init....")
	structLiteralInit()

	fmt.Println("********************** Examples with Array ***********************")
	initArray()
	fmt.Println("Slice an array:")
	initSlice()
	slice()
	sliceByMake()
	appendSlice()
	useRange()
}
