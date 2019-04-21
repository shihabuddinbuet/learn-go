package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("summation : ", sum)

	sum = 1
	// "while" loop using "for" loop
	for sum < 100 {
		sum += sum
	}
	fmt.Println("sum : ", sum)

	// infinity loop
	for {

	}
}
