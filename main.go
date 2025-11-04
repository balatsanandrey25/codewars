package main

import "fmt"

func main() {
	num := 123532
	fmt.Println(isPalindrome(num))
}
func isPalindrome(x int) bool {
	// Проверить отрицательные числа - если число отрицательное, сразу вернуть false
	if x < 0 {
		return false
	}
	// Создаем реверсивное чило
	rev := reverseHalfNumber(x)
	// Сравноваем иходное число с реверсивным
	return rev == x
}

func reverseHalfNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}
