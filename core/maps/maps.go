package main

import "fmt"

func iterateMap() {
	// the order of traversal is not certain
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"China":  "Beijing",
		"Japan":  "Tokyo",
	}

	// starting with Go 1.12, the fmt package prints maps in key-sorted order to ease testing
	// but not in loop
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func mapSlice() {
	// create a slice whose element type is a map
	items := make([]map[string]int, 5)
	for i := range items {
		items[i] = make(map[string]int, 2)
		items[i]["one"] = 1
		items[i]["two"] = 2
	}

	fmt.Printf("Value of items: %v\n", items)
}

func main() {
	iterateMap()
	mapSlice()
}
