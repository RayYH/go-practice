package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestMapDeclaration(t *testing.T) {
	// declaration
	var mapList map[string]int
	var mapAssigned map[string]int
	// initialization
	mapList = map[string]int{
		"one": 1,
		"two": 2,
	}
	// be assigned a exists map
	mapAssigned = mapList

	assert.Equal(t, mapList, map[string]int{
		"one": 1,
		"two": 2,
	})
	assert.Equal(t, mapList["one"], 1)
	assert.Equal(t, mapAssigned["two"], 2)

	// reference type can be declared via make keyword, DO NOT use new keyword
	mapCreated := make(map[string]float64)
	// dynamically add items
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	// modify items
	mapAssigned["two"] = 3

	assert.Equal(t, mapAssigned["two"], 3)
	assert.Equal(t, mapCreated["key2"], 3.14159)
}

func TestMapValueCanBeAnyType(t *testing.T) {
	mf := map[int]func() int{
		1: func() int {
			return 1
		},
		2: func() int {
			return 2
		},
		3: func() int {
			return 3
		},
	}

	for k, v := range mf {
		assert.Equal(t, v(), k)
	}
}

func TestMapElement(t *testing.T) {
	var value int
	var isPresent bool
	mapList := make(map[string]int)
	mapList["Beijing"] = 100
	mapList["Shanghai"] = 99
	mapList["Nanjing"] = 89

	// exists item
	value, isPresent = mapList["Beijing"]
	assert.Equal(t, value, 100)
	assert.True(t, isPresent)

	// non-exists item
	value, isPresent = mapList["Paris"]
	assert.Equal(t, value, 0) // default int value
	assert.False(t, isPresent)

	// delete map item
	delete(mapList, "Beijing")
	value, isPresent = mapList["Beijing"]
	assert.Equal(t, value, 0) // default int value
	assert.False(t, isPresent)
}

func ExampleSortMap() {
	// A map is an unordered collection of key-value pairs.
	// If you need a stable iteration order, you must maintain a separate data structure.
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	// Output:
	// Alice 23
	// Bob 25
	// Eve 2
}

func ExampleInvertMap() {
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}
	invMap := make(map[int]string)

	for k, v := range m {
		invMap[v] = k
	}

	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	// Output:
	// Key: 25, Value: Bob
	// Key: 23, Value: Alice
	// Key: 2, Value: Eve
}
