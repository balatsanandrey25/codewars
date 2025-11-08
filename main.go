package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	in := make(chan int)
	go func() {
		for i := range 10 {
			in <- i + 1
		}
		close(in)
	}()
	return in
}
func double(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range ch {
			out <- v * 2
			time.Sleep(1000 * time.Millisecond)
		}
		close(out)
	}()

	return out
}
func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	reader(double(writer()))
}
