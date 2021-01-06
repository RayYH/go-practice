package maps

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBasicOperationsOnMaps(t *testing.T) {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	// The make built-in function allocates and initializes an object
	persons := make(map[string]PersonInfo)
	// add string-PersonInfo pairs
	persons["1"] = PersonInfo{"1", "Ray", "Street 1"}
	persons["2"] = PersonInfo{"2", "Tom", "Street 2"}
	// fetch element
	person, ok := persons["1"]
	assert.Equal(t, PersonInfo{"1", "Ray", "Street 1"}, person)
	assert.True(t, ok)
	// remove element
	delete(persons, "1")
	_, ok = persons["1"]
	assert.False(t, ok)
}

func TestDeclarationsOfMaps(t *testing.T) {
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
	assert.Equal(t, mapList, mapAssigned)

	// reference type can be declared via make keyword, DO NOT use new keyword
	mapCreated := make(map[string]float64)
	// dynamically add items
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	assert.Equal(t, mapCreated["key2"], 3.14159)
}

func TestValueOfMapsCanBeAnyType(t *testing.T) {
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

func TestFetchingElementsFromMaps(t *testing.T) {
	var value int
	var isPresent bool
	// use make to create a map
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

func TestInvertingMaps(t *testing.T) {
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}
	invMap := make(map[int]string)

	for k, v := range m {
		invMap[v] = k
	}

	assert.Equal(t, invMap, map[int]string{
		23: "Alice",
		2:  "Eve",
		25: "Bob",
	})
}

func TestElementsOfSlicesCanBeMapType(t *testing.T) {
	// items is a slice whose elements are of map[string]int type
	items := make([]map[string]int, 2)
	for i := range items {
		items[i] = make(map[string]int, 3)
		items[i]["one"] = 1
		items[i]["two"] = 2
	}

	assert.Equal(t, items, []map[string]int{{"one": 1, "two": 2}, {"one": 1, "two": 2}})
}

func TestTheOrdersOfIteratingElementsOfMapAreNotCertain(t *testing.T) {
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
		assert.NotNil(t, key)
		assert.NotNil(t, capitals[key])
	}
}

func TestAStableWayToSortMaps(t *testing.T) {
	// A map is an unordered collection of key-value pairs.
	// If you need a stable iteration order, you must maintain a separate data structure.
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	values := []int{23, 25, 2}
	index := 0

	for _, k := range keys {
		assert.Equal(t, m[k], values[index])
		index++
	}
}
