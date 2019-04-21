package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float32
}

/*
	Decalre a map: map[keyType]ValueType
*/
func mapByMake() {
	m := make(map[string]Vertex)
	m["Dhaka"] = Vertex{90.1, 80.4}
	fmt.Println(m["Dhaka"])
}

/*
	Initialization of Map is like the initalization of array of slice
	map[keyType]valueType {"key":"value"}
*/
func mapByLiteral() {
	cityMap := map[string]Vertex{
		"Dhaka": Vertex{1, 2},
		"Delhi": Vertex{3, 4},
	}
	fmt.Println("cityMap: ", cityMap)
}

func mutateMap() {
	m := make(map[string]int)
	fmt.Println("Insert inside map")
	m["dhaka"] = 1
	m["tokyo"] = 2
	fmt.Println("Map content after insertion : ", m)
	fmt.Println("Dhaka: ", m["dhaka"], "tokyo:", m["tokyo"])

	m["dhaka"] = 4
	fmt.Println("Map content after update : ", m)

	delete(m, "tokyo")
	fmt.Println("Map content after removing tokyo : ", m)
}

func main() {
	mapByMake()
	mapByLiteral()
	mutateMap()
}
