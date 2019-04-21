package main

import (
	"fmt"
)

//***********************simple function *******************************

/*while declaring a variable, variable type comes after the variable name
For example: x int, here "x" is variable and "int" is data type
*/

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

//************************* function continued ***************************

/*
When two or more consecutive parameter has same data type (like the above functions),
the parameters declaration can be shortened by omiting the type from each parameter but
the last one.
*/

func addThreeInteger(x, y, z int) int {
	return x + y + z
}

func subtractThreeInteger(x, y, z int) int {
	return x - y - z
}

//**************************** return multiple results ****************
func swap(x, y string) (string, string) {
	return y, x
}

//*********************** Named return ********************
func fmtSum(x, y int) (message string, sum int) {
	return "sum of two integer: ", x + y
}

func main() {
	fmt.Println("Hello world!!!!")
	fmt.Println("20 + 15 =", add(20, 15))
	fmt.Println("20 - 15 =", subtract(20, 15))

	fmt.Println("10 + 2 + 5 = ", addThreeInteger(10, 2, 5))
	fmt.Println("10 - 2 - 5 = ", subtractThreeInteger(10, 2, 5))

	a, b := swap("shihab", "uddin")
	fmt.Println("a = ", a, "b=", b)

	fmt.Println(fmtSum(10, 12))

	fmt.Println("****************** function by value  ***********")
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println("Compute : ", compute(add, 20))
	adder1 := adder()
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = adder1(i)
	}
	fmt.Println("summation using funcation by value: ", sum)
}
