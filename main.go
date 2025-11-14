package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	// mu := sync.Mutex{}
	var counter int64 = 0
	// or
	// counter := 0
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// defer mu.Unlock()
			// mu.Lock()
			atomic.AddInt64(&counter, 1)
			// counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
