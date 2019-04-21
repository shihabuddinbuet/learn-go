package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello world from init function")
}
func main() {
	fmt.Println("Hello world from main function !!!!")
	add(1, 2)
}
