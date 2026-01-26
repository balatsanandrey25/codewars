package main

import "fmt"

func DigPow(n, p int) int {
	// your code (ap+bp+1+cp+2+dp+3+...)=n∗k
	str := fmt.Sprintf("%d", n)
	result := 0.0
	for i, char := range str {
		result += float64(powInt(int(char-'0'), p+i))
	}
	k := result / float64(n)
	if k == float64(int(k)) {
		return int(k)
	} else {
		return -1
	}
}
func powInt(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
func main() {
	fmt.Println(DigPow(46288, 3))
}
