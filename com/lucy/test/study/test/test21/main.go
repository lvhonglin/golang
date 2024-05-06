package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	println(ctx.Done() == nil)
	println(cancelFunc)
	context.WithTimeout(context.Background(), time.Second)

}
