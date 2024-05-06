package main

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		go func() {
			time.Sleep(time.Second)
			println("789")
		}()
		println(123)
		runtime.Goexit()
		println(456)
	}()

	time.Sleep(time.Second * 5)
	println("over")
}
