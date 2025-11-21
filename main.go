package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charCount := make(map[rune]int)
	// Заполняем map символами из первой строки
	for _, char := range s {
		charCount[char]++
	}
	// Проверяем вторую строку
	for _, char := range t {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "jam" //r2 a2 c2 e1
	t := "jar" //
	fmt.Println(isAnagram(s, t))
}
