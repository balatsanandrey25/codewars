package main

import "fmt"

func FindOdd(seq []int) int {
	counts := make(map[int]int)
	// Множество чисел с нечетным количеством
	oddNumbers := make(map[int]bool)

	for _, num := range seq {
		counts[num]++
		if counts[num]%2 != 0 {
			oddNumbers[num] = true
		} else {
			delete(oddNumbers, num)
		}
	}

	// Найти максимум в oddNumbers
	maxOdd := -1
	for num := range oddNumbers {
		if num > maxOdd {
			maxOdd = num
		}
	}

	return maxOdd
}

func main() {
	arr := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5, 7, 7, 99}
	result := FindOdd(arr)
	fmt.Printf("Максимальное число с нечетным количеством повторений: %d\n", result)
	FindOdd(arr)
}
