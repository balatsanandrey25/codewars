package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(ctx context.Context, val int) int {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()
	select {
	case <-ch:
		return val * 2
	case <-ctx.Done():
		return -1
	}
}

func main() {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	in := make(chan int)
	out := make(chan int)

	go func() {
		defer close(in)
		for i := range 10 {
			in <- i + 1
		}

		select {
		case <-in:
		case <-ctx.Done():
			return
		}
	}()
	now := time.Now()
	processParallel(ctx, in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

// Операция не должна выполняться не более 5 сек
func processParallel(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
	wg := sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				select {
				case out <- processData(ctx, v):
				case <-ctx.Done():
					return
				}
			}
		}()

	}
	go func() {
		wg.Wait()
		close(out)
	}()
}
