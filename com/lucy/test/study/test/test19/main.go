package main

import (
	"context"
	"time"
)

func main() {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	time.Sleep(time.Second * 3)
	println(1)
	go func() {
		<-timeout.Done()
		println("over1")
		println(timeout.Err().Error())

	}()
	time.Sleep(time.Second * 3)
	go func() {
		<-timeout.Done()
		println("over2")
		println(timeout.Err().Error())
	}()
	time.Sleep(time.Second * 10)
	cancelFunc()
	println("end")
}
