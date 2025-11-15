package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func unpredictableFunc() int {
	n := rand.Intn(40)
	fmt.Println(n)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}
func predictableFunc(ctx context.Context) (int, error) {
	c, _ := context.WithTimeout(ctx, 5*time.Second)
	ch := make(chan struct{})
	var result int
	go func() {
		result = unpredictableFunc()
		close(ch)
	}()
	select {
	case <-c.Done():
		return -1, errors.New("Функция работает более 5 секунд")
	case <-ch:
		return result, nil
	}

}
func main() {
	v, e := predictableFunc(context.Background())
	fmt.Println(v, e)
}
