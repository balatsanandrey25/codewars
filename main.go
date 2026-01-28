package main

import "fmt"

func topKFrequent(nums []int, k int) []int {

	collectionNums := make(map[int]int, len(nums))
	var resultNums []int
	for _, v := range nums {
		collectionNums[v] += 1
	}
	for v, key := range collectionNums {
		if key >= k {
			resultNums = append(resultNums, v)
		}
	}
	if resultNums != nil {
		return resultNums
	} else {
		return nums
	}
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
