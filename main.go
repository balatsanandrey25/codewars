package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	searchTarget := make(map[int]int, 0)
	var difference int
	for i, v := range nums {
		difference = target - v
		if difference < 0 {
			continue
		}
		valueMap, ok := searchTarget[difference]
		if ok {
			return []int{valueMap, i}
		}
		searchTarget[v] = i

	}
	return []int{0, 0}
}
