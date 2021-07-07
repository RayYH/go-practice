package advanced_types

import (
	"log"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarationsAndInitializationsOfMaps(t *testing.T) {
	t.Run("non-initialized map is nil", func(t *testing.T) {
		var mapList map[string]int
		assert.Nil(t, mapList)
	})

	t.Run("maps are referenced types", func(t *testing.T) {
		mapList := map[string]int{
			"one": 1,
			"two": 2,
		}
		mapDuplicated := mapList
		mapDuplicated["one"] = 3
		assert.Equal(t, mapList["one"], 3)
	})

	t.Run("using make to create maps", func(t *testing.T) {
		mapCreated := make(map[string]float64)
		mapCreated["key1"] = 4.5
		mapCreated["key2"] = 3.14159
		assert.Equal(t, mapCreated["key2"], 3.14159)

		// never use new() to create a map
		mapDuplicated := new(map[string]float64)
		assert.Panics(t, func() {
			(*mapDuplicated)["self"] = 1.1
		})

		mapDuplicated = &mapCreated
		assert.NotNil(t, mapDuplicated)

		(*mapDuplicated)["key1"] = 3.5
		assert.Equal(t, mapCreated["key1"], 3.5)
	})
}

func TestCURDOnMaps(t *testing.T) {
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}
	// declare a map
	persons := make(map[string]PersonInfo)

	// add element (k/v)
	persons["1"] = PersonInfo{"1", "Ray", "Street 1"}
	persons["2"] = PersonInfo{"2", "Tom", "Street 2"}

	// get v by k
	person, ok := persons["1"]
	assert.Equal(t, PersonInfo{"1", "Ray", "Street 1"}, person)
	assert.True(t, ok)

	// modify v by k
	person.Name = "Rayyh"
	persons["1"] = person
	person, ok = persons["1"]
	assert.Equal(t, PersonInfo{"1", "Rayyh", "Street 1"}, person)
	assert.True(t, ok)

	// remove element by k
	delete(persons, "1")
	_, ok = persons["1"]
	assert.False(t, ok)
}

func TestMapsWithOtherDataTypes(t *testing.T) {
	t.Run("func type", func(t *testing.T) {
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
	})

	t.Run("slice type", func(t *testing.T) {
		data := make(map[int][]int)
		data = map[int][]int{
			1: {1, 2},
			2: {2, 3},
		}
		assert.Equal(t, data[1][0], 1)
	})

	t.Run("element of slice can be of map type", func(t *testing.T) {
		items := make([]map[string]int, 2)
		for i := range items {
			items[i] = make(map[string]int, 3)
			items[i]["one"] = 1
			items[i]["two"] = 2
		}

		assert.Equal(t, items, []map[string]int{{"one": 1, "two": 2}, {"one": 1, "two": 2}})
	})
}

func TestMapOperations(t *testing.T) {
	t.Run("inverse map", func(t *testing.T) {
		m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25, "Mary": 25}
		invMap := make(map[int]string)

		for k, v := range m {
			invMap[v] = k
		}

		// since iteration order is not certain, so value of key 25 can be "Bob" or "Mary"
		assert.Equal(t, invMap[23], "Alice")
		assert.Equal(t, invMap[2], "Eve")
	})

	t.Run("the orders of iterating elements of maps are not certain", func(t *testing.T) {
		// the order of traversal is not certain
		capitals := map[string]string{
			"France": "Paris",
			"Italy":  "Rome",
			"Japan":  "Tokyo",
			"China":  "Beijing",
		}

		// starting with Go 1.12, the fmt package prints maps in key-sorted order to ease testing
		// but not in loop
		for key := range capitals {
			log.Println(key)
			assert.NotNil(t, key)
			assert.NotNil(t, capitals[key])
		}
	})

	t.Run("a stable way to sort maps", func(t *testing.T) {
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
	})
}
