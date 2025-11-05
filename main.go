package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func newNode(data int) *Node {
	return &Node{data: data, next: nil}
}
func (ll *LinkedList) append(data int) {
	newNode := newNode(data)

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (ll *LinkedList) display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	list6 := &LinkedList{
		head: &Node{
			data: 1,
			next: &Node{data: 0, next: &Node{data: 0, next: &Node{data: 1,
				next: nil}}}},
	}

	// Второй такой же список
	list7 := &LinkedList{
		head: &Node{
			data: 1,
			next: &Node{data: 2, next: &Node{data: 2, next: &Node{data: 0,
				next: nil}}}},
	}
	addTwoNumbers(list6, list7).display()
}
func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	i1, err := collectorNums(l1)
	if err != nil {
		return &LinkedList{
			head: &Node{
				data: 0,
				next: nil,
			}}
	}
	i2, err := collectorNums(l2)
	if err != nil {
		return &LinkedList{
			head: &Node{
				data: 0,
				next: nil,
			}}
	}

	totalSumStr := strconv.Itoa(i1 + i2)
	newLinkedList := &LinkedList{}
	for _, v := range totalSumStr {
		tempInt, _ := strconv.Atoi(string(v))
		newLinkedList.append(tempInt)
	}

	return newLinkedList
}
func collectorNums(list *LinkedList) (int, error) {
	result := make([]int, 0)
	current := list.head
	for current.next != nil {
		result = append(result, current.data)
		current = current.next
	}
	result = append(result, current.data)
	var buffer bytes.Buffer
	for _, num := range result {
		buffer.WriteString(strconv.Itoa(num))
	}

	str := buffer.String()
	resInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return resInt, nil
}
