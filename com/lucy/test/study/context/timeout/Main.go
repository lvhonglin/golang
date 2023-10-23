package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func coroutine(ctx context.Context, duration time.Duration, id int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("协程%d退出\n", id)
			wg.Done()
			return
		case <-time.After(duration):
			fmt.Printf("消息来自协程%d\n", id)
		}
	}
}
func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go coroutine(ctx, 1*time.Second, i, wg)
	}
	wg.Wait()
	fmt.Printf("结束%v", "了")
}
