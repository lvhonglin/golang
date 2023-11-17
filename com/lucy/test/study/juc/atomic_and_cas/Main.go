package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value atomic.Value
	value.Store(1)
	val := value.Load()
	real := val.(int)
	fmt.Println(real)
	swapped := value.CompareAndSwap(2, 2)
	fmt.Println(swapped)
	println("=================")
	var a int32 = 10
	atomic.AddInt32(&a, 1)
	atomic.AddInt32(&a, 1)
	println(a)
}
