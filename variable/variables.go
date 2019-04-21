package main

import (
	"fmt"
)

/*
var statement declare list of variable as parameters in function
where type is last.
For details, multiple consecutive same type variables can be shortened as
only the variable names but the last one. Last variable declaration will contain
both varibable name and type
*/

var java, python, golang bool

//************************* constant ***********************
/*
	Like the var const can be package level and function level
	while decalring constant, use = to assign value
*/

const (
	PI   = 3.14
	SQRT = 1.414
	PATH = "//"
)

func main() {
	var i int
	fmt.Println(i, java, python, golang)

	//****************************** variable with initailization *****************
	/*
		while initalizition varibales at the time of declaration, type can be omitted.
		Providing type at initialization is optional
	*/

	var x, y, z int = 1, 2, 3
	var a, b, c = true, "shihab", 1.4
	fmt.Println(x, y, z)
	fmt.Println(a, b, c)

	/*
		Inside function the variable initialization can be shortened using ":=" without var statement.
	*/

	k := 10
	fmt.Println("k := ", k)

	//some basic types
	var (
		name   string = "shihab"
		knowGo bool   = false
	)
	fmt.Printf("Type : %T, Value :%v\n", name, name)
	fmt.Printf("Type: %T, Value: %v\n", knowGo, knowGo)

	//zero values
	/*
		After variable decalartion it is initialized with default values.
		For int,float32,bool the zero values are 0,0,false respectively
	*/

	var m int
	var n float32
	var p bool

	fmt.Println("Int,Float32,bool zero values: ", m, n, p)

	//***************************** Type conversion **************************
	//T(v) expression converts value v to type T

	ft := float32(m)
	fmt.Printf("Type of ft: %T and value : %v\n", ft, ft)

}

/*
var statement can be package level as well as function level.
Here java,python,golang variables are in package level. On the other hand variable i
is in function level
*/
