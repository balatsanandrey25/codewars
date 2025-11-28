package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// Преобразуем строку в срез рун для сортировки
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)

		// Добавляем исходную строку в группу анаграмм
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
		fmt.Println("sortedStr:", sortedStr, ";", str)
	}

	// Преобразуем map в срез срезов
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
func main() {
	// Тестовые примеры
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"stop", "pots", "tops", "opts", "post"},
	}

	for _, strs := range testCases {
		_ = groupAnagrams(strs)
	}
}
