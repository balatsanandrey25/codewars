package main

import "fmt"

func main() {
	romNums := "III"
	fmt.Println(romanToInt(romNums))
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	} // Создали мару для хранения римских чисел
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		current, ok := romanMap[s[i]]
		if !ok {
			return 0
		}
		if i < len(s)-1 && current < romanMap[s[i+1]] {
			result -= current
		} else {
			result += current
		}
	}
	return result
}
