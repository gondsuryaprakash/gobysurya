package maps

import "fmt"

func CreatingMaps() {

	maps := make(map[string]int)
	// Print empty maps

	fmt.Println("Maps : ", maps)

	// Adding key/value pair in maps

	maps["a"] = 1
	maps["b"] = 2
	maps["c"] = 3

	fmt.Println("Maps :", maps)

}
