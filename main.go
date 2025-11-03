package main

import (
	"fmt"
)

func main() {
	busStops := [][2]int{
		{3, 0}, // На остановку зашли 3, вышли 0
		{2, 1}, // На остановку зашли 2, вышли 1
		{1, 2}, // На остановку зашли 1, вышли 2
	}
	fmt.Println(Number(busStops))
}
func Number(busStops [][2]int) int {
	passagire := 0
	for _, stop := range busStops {
		passagire += stop[0] - stop[1]
	}

	return passagire
}
