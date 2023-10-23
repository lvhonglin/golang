package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("任务%v结束退出\n", ctx.Value(key))
			return
		default:
			fmt.Printf("任务%v正在运行中\n", ctx.Value(key))
			time.Sleep(time.Second * 2)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, key, "监控1")
	go run(ctx)
	time.Sleep(time.Second * 10)
	fmt.Println("停止任务")
	cancel()
	time.Sleep(time.Second * 3)
}
