package main

// it take func as value
func compute(add func(int, int) int, a int) int {
	return add(add(3, 2), a)
}

// it return func as value
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
