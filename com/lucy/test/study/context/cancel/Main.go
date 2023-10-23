package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("任务%v结束退出\n", id)
			return
		default:
			fmt.Printf("任务%v正在运行中\n", id)
			time.Sleep(time.Second * 2)
		}

	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go run(ctx, 1)
	go run(ctx, 2)
	time.Sleep(time.Second * 10)
	fmt.Println("停止任务1")
	cancel()
	time.Sleep(time.Second * 3)
	return
}
