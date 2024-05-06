package main

import (
	"sync"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("rrrr")
		}
	}()
	//goroutineCount := 3
	taskChan := make(chan int, 1213)
	wait := sync.WaitGroup{}
	//for i := 0; i < goroutineCount; i++ {
	//	go func() {
	//		defer func() {
	//			if r:=recover();r!=nil{
	//				println(r)
	//			}
	//		}()
	//		for a := range taskChan {
	//			println(a)
	//			wait.Done()
	//			panic("www")
	//		}
	//	}()
	//}
	for i := 0; i < 1000; i++ {
		wait.Add(1)
		taskChan <- i
	}
	wait.Wait()
}
