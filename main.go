package main

import (
	"fmt"
	"sort"
)

func main() {
	ln1 := []int{1, 2, 3, 4}
	ln2 := []int{1}
	fmt.Println(findMedianSortedArrays(ln1, ln2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return float64(0)
	}
	if len(nums2) == 0 {

		return getMidle(nums1)
	}
	if len(nums1) == 0 {

		return getMidle(nums2)
	}

	p1, n1 := nums1[0], nums1[len(nums1)-1]
	p2, n2 := nums2[0], nums2[len(nums2)-1]

	if n1 <= p2 {
		return getMidle(append(nums1, nums2...))
	}
	if n2 <= p1 {
		return getMidle(append(nums2, nums1...))
	}

	sliceRes := append(nums1, nums2...)
	sort.Ints(sliceRes)
	return getMidle(sliceRes)
}

// Функция выдает середину слайса
func getMidle(slice []int) float64 {
	lenS := len(slice)
	if lenS == 1 {
		return float64(slice[0])
	}
	mid := lenS / 2
	if lenS%2 == 0 {
		return float64(slice[mid-1]+slice[mid]) / 2.0
	}
	return float64(slice[mid])
}
