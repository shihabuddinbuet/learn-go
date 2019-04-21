package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Atoi(s string) {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Can not convert ", s, "to int")
		return
	}
	fmt.Println("integer value : ", i)
}

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() float64 {
	return float64(e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorNegativeSqrt(x).Error()
	}
	return 0, nil
}

func main() {
	if error := run(); error != nil {
		fmt.Println(error)
	}
	Atoi("12")
	Atoi("ab")
}
