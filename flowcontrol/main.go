package main

import (
	"fmt"
)

func greater(a, b int) bool {
	if a > b {
		return true
	}
	return false
	// this can be written in a single line.
	// just "return a > b"
}

func greaterOrEqual(a, b int) bool {
	if a > b || a == b {
		return true
	}
	return false
	// this can be written as "return a > b || a == b"
}

//******************* if in short version *******************
const (
	SUMLIMIT = 1000
)

/*
	Like for, the if statement can start with a short statment to execute
	if the condition is statisfied
*/
func sum(a, b int) int {
	if v := a + b; v < SUMLIMIT {
		return v
	}
	return SUMLIMIT
}

//******************** if else example with short version ********************

func sumWithMessage(a, b int) (string, int) {
	if v := a + b; v < SUMLIMIT {
		return "Summation is within limit", v
	} else {
		fmt.Printf("Summation = %v is exceed limit = %v . Returing the summation limit", v, SUMLIMIT)
	}
	return "Summation execeed the limit", SUMLIMIT
}

func testSwitch() {
	name := "shihab"
	fmt.Println(nameExits(name))
	showWhenStaruday()
}

func sayGoodBye() string {
	return "GOODBYE. SEE YOU AGAIN"
}

func main() {
	/*
		Defer statement enforces to defer the execution of the function until
		the surrounding function called is executed.
		Defer statement is evaluated but executed after surrounding functions execution
		is finished.

		If there is multiples defer statements, the statements are executeed in LIFO (last in first out)
		order..
	*/

	defer fmt.Println(sayGoodBye() + " AGAIN")
	defer fmt.Println(sayGoodBye())

	var a, b = 10, 10
	fmt.Println(greater(a, b))
	fmt.Println(greaterOrEqual(a, b))
	fmt.Println("summation with if short version = ", sum(10, 20))
	message, sum1 := sumWithMessage(10, 20)
	fmt.Println(message, sum1)

	testSwitch()
}
