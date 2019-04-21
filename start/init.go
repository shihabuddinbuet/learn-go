package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello world from init.go ")
}

func add(x, y int) {
	fmt.Println(x + y)
}
