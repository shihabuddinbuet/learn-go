package main

import (
	"fmt"
)

/*
	if type is not matched, return default value and false
	if boolean value is not received, it get painc if assertion fail
*/

func typeAssert(value interface{}) {
	v, ok := value.(string)

	if ok {
		fmt.Println("It's string and value : ", v)
		return
	}

	i, ok := value.(int)
	if ok {
		fmt.Println("It's int and value : ", i)
		return
	}

}

func typeAssertBySwitch(i interface{}) {

	switch v := i.(type) {
	case string:
		fmt.Println("It's string, value = ", v)
	case int:
		fmt.Println("It's int, value = ", v)
	default:
		fmt.Println("it's neither string nor int, value = ", v)
	}
}
