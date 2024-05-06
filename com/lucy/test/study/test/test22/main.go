package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Parent struct {
	a int
}

func (p *Parent) test() {
	println(13)
	println(p.a)
}
func main() {
	var counter atomic.Value
	counter.Store(0)

	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&counter, &wg)
	go increment(&counter, &wg)

	wg.Wait()

	fmt.Println("Counter:", counter.Load().(int)) // 输出的counter值应该是预期的结果
}

func increment(counter *atomic.Value, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		for {
			value := counter.Load().(int)
			swapped := counter.CompareAndSwap(value, value+1)
			if swapped {
				break
			}
		}
	}
	wg.Done()
}
