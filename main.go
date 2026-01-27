package main

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!
	str := string(s)

	return str == strings.ToUpper(str)

}
func main() {

}
