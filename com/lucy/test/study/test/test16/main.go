package main

import "time"

func main() {
	var a chan int
	go func() {
		println("进来了")
		<-a
		println("出来了")
	}()
	time.Sleep(time.Second * 5)
}
