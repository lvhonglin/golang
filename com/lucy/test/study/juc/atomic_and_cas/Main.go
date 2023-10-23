package main

import (
	"fmt"
	"go.uber.org/atomic"
)

func main() {
	var value atomic.Value
	value.Store(1)
	val := value.Load()
	real := val.(int)
	fmt.Println(real)
	swapped := value.CompareAndSwap(2, 2)
	fmt.Println(swapped)
}
