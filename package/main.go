package main

//Every programme start running from "main" package

/*
	if the files start with other than main package for example package abc, It will get the
	following error while running the file:
	go run: cannot run non-main package
*/

import (
	"fmt"
	"math/rand"
)

/*
A package comprises of files that start with the package name.
For example "fmt" package comprises of the files that start with "package fmt".
*/

func main() {
	fmt.Println("Hello world!!!!")
	fmt.Println("", rand.Intn(10))
}
