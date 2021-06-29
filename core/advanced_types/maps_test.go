package advanced_types

import (
	"log"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// map 的声明与初始化
func TestDeclarationsAndInitializationsOfMaps(t *testing.T) {
	t.Run("non-initialized map is nil", func(t *testing.T) {
		// map 是可以动态增长的，因此声明时不需要知道 map 的长度
		var mapList map[string]int
		var mapAssigned map[string]int
		// 未初始化的 map 的值是 nil
		assert.Nil(t, mapList)
		assert.Nil(t, mapAssigned)
	})

	t.Run("maps are referenced types", func(t *testing.T) {
		// 直接通过字面量 ({k1:v1, k2:v2} 风格) 来进行初始化，这里没有显式声明变量类型，因为编译器可以自行推断
		mapList := map[string]int{
			"one": 1,
			"two": 2,
		}
		// 从已有的 map 直接赋值
		mapDuplicated := mapList

		// mapDuplicated 没有经过深度拷贝，因此修改它会同时修改原有的 mapList
		mapDuplicated["one"] = 3
		assert.Equal(t, mapList["one"], 3)
	})

	// 我们可以使用 `make` 方法来初始化 map 为零值
	// 而 `new` 方法返回的是一个内存经过初始化的指针 (永远不要使用 `new` 来初始化一个 map)
	t.Run("using make to create maps", func(t *testing.T) {
		mapCreated := make(map[string]float64)
		mapCreated["key1"] = 4.5
		mapCreated["key2"] = 3.14159
		assert.Equal(t, mapCreated["key2"], 3.14159)

		// 这里 mapDuplicated 是一个指向 map[string]float64 类型的指针
		mapDuplicated := new(map[string]float64)
		// 在这里 map 所需要的内存是 0，因此这里的 new 方法创建的内存实际上是 Nil，往其中直接追加元素是非法的
		assert.Panics(t, func() {
			(*mapDuplicated)["self"] = 1.1
		})

		// 但是可以将一块已经初始化内存的 map[string]float64 赋给 mapDuplicated
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
	// 使用 make 初始化一个 map
	persons := make(map[string]PersonInfo)

	// 添加 k/v
	persons["1"] = PersonInfo{"1", "Ray", "Street 1"}
	persons["2"] = PersonInfo{"2", "Tom", "Street 2"}

	// 根据 k 获取 v
	person, ok := persons["1"]
	assert.Equal(t, PersonInfo{"1", "Ray", "Street 1"}, person)
	assert.True(t, ok)

	// 修改 v 需要根据 k 重新赋值
	person.Name = "Rayyh"
	persons["1"] = person
	person, ok = persons["1"]
	assert.Equal(t, PersonInfo{"1", "Rayyh", "Street 1"}, person)
	assert.True(t, ok)

	// 根据 k 移除元素
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

		// 直接 k/v 对调即可，注意如果 v 相同，会被替换掉
		for k, v := range m {
			invMap[v] = k
		}

		// 由于遍历的顺序是不确定的，因此你无法判断处 25 对应的值是 "Bob" 还是 "Mary"
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
			// 如果使用 log 包的话，这里每一次打印的 key 的顺序都不一致
			// 但是如果使用 fmt.Println() 的话，则是按 key 顺序排列的
			// 综上，map 的遍历顺序是不能够被确定的
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

		// 让 map 遍历的顺序稳定的方法就是对 k 进行排序，然后对 k 进行遍历
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
