package main

func towSum(nums []int, target int) []int {
	cache := map[int]int{}
	for i, v := range nums {
		if value, exists := cache[target-v]; exists {
			return []int{value, i}
		}
		cache[v] = i
	}

	return []int{}
}

func main() {

}
