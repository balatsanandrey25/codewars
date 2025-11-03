package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(Persistence(0))
}

func Persistence(n int) int {
	// Базовый случай: если число уже однозначное
	if n < 10 {
		return 0
	}

	counter := 0
	current := n

	for current >= 10 {
		current = multiplyDigits(current)
		counter++
	}

	return counter
}

func multiplyDigits(n int) int {
	str := strconv.Itoa(n)
	product := 1

	for i := 0; i < len(str); i++ {
		digit, _ := strconv.Atoi(string(str[i]))
		product *= digit
	}

	return product
}
