package main

import (
	"fmt"
	"time"
)

const (
	SHIHAB = "shihab"
	UDDIN  = "uddin"
	HASAN  = "hasan"
)

/*
	switch in go is similar of c,c++ execept that case evaluation happend from top
	to bottom and stop when any case is satisfied
*/

func nameExits(name string) bool {
	switch name {
	case SHIHAB:
		return true
	case UDDIN:
		return true
	case HASAN:
		return true
	default:
		return false
	}
}

/*
	This is shorter version of writing switch, like "for" loop or "if" condition
*/

func showWhenStaruday() {

	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	default:
		fmt.Println("Far away")
	}
}

// This switch without condition
func greet() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
