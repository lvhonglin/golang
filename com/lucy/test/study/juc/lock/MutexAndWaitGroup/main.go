package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var num int32
var operatorMux sync.Mutex
var wg sync.WaitGroup

func init() {
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			print_lock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func print_lock() {
	operatorMux.Lock()
	println(num)
	atomic.AddInt32(&num, 1)
	time.Sleep(time.Second)
	operatorMux.Unlock()
}
