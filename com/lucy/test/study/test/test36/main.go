package main

import (
	"sync"
	"time"
)

type MyConcurrentMap struct {
	sync.Mutex
	mp map[int]int
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp: make(map[int]int),
	}
}

var l sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		lock := l
		lock.Lock()
		println("111111")
		time.Sleep(time.Second * 5)
		lock.Unlock()
		wg.Done()
	}()
	go func() {
		lock := l
		lock.Lock()
		println("222222")
		time.Sleep(time.Second * 5)
		lock.Unlock()
		wg.Done()
	}()
	wg.Wait()

}
