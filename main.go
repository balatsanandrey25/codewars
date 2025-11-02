package main

import (
	"fmt"
)

func main() {

	fmt.Println(MultiTable(5))
}
func MultiTable(number int) string {
	//good luck
	result := ""
	for i := 1; i <= 10; i++ {
		if i < 10 {
			result += fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
		} else {
			result += fmt.Sprintf("%d * %d = %d", i, number, i*number)
		}
	}
	return result

}
