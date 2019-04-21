package main

import (
	"fmt"
)

/*
	Declaring array is simlar to declaring a variable.
	var <variable_name> <type>
*/

func initArray() {
	var names [2]string
	names[0] = "shihab"
	names[1] = "uddin"
	fmt.Println("first and second names", names[0], names[1])

	numbers := [3]int{1, 2, 3}
	fmt.Println("numbers: ", numbers)

}

type point struct {
	X int
	Y int
}

/*
	Slice is like an array without size
*/

func initSlice() {
	points := []point{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println("points: ", points)

}

/*
	format to take a slice is array[low:high]
	where low is inclusive and high is exclusive index.
	Default value of both low and length of underlying array is high bound
	For example , a := b[0:], a and b will same elements

	Slice is like referencing. Any value is changed, affect all reference value.
	Changing b[0] means change of names[1] and it changes a[1] also as a refer to names.

*/

func slice() {
	numbers := [5]int{1, 2, 3, 4, 5}
	var snumbers []int = numbers[1:3]
	fmt.Println("numbers: ", numbers)
	fmt.Println("sliced numbers: ", snumbers)
	fmt.Println("slice numbers with slice default: ", numbers[:3])

	fmt.Println("slice reference....")
	names := [4]string{"shihab", "uddin", "mike", "tom"}
	fmt.Println("names: ", names)
	a := names[0:3]
	b := names[1:4]
	fmt.Println("a: ", a, "b:", b)
	b[0] = "xxx"
	fmt.Println("names: ", names)
	fmt.Println("a: ", a, "b:", b)
}

func sliceByMake() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
}

// append is possible with nill slice

func appendSlice() {
	var s []int
	s = append(s, 0)
	printSlice("0 append", s)
	s = append(s, 1)
	printSlice("1 append", s)
	s = append(s, 1, 2, 3)
}

func useRange() {
	s := []int{1, 2, 3, 4}
	for i, v := range s {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	for _, v := range s {
		fmt.Println("value: ", v)
	}
}
func printSlice(str string, s []int) {
	fmt.Printf("%s, len=%d, cap=%d,%v\n", str, len(s), cap(s), s)
}
