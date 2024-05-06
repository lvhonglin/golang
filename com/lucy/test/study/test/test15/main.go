package main

import (
	"context"
	"sync"
	"time"
)

// 实现一个map
// 面向高并发
// 插入和查询时间复杂度是O(1)
// 查询时如果key存在则返回val；若key不存在则阻塞到key val对被放入后，获取val返回
// 写出真实代码，不能有死锁或者panic风险
func main() {
	syncMap := NewSyncMap()
	go func() {
		println("读取协程，开始阻塞")
		get, err := syncMap.Get("test")
		if err != nil {
			println("err")
		}
		println("读取协程，获取到结果:" + get)
	}()
	go func() {
		time.Sleep(time.Second * 10)
		syncMap.put("test", "123")
	}()
	time.Sleep(time.Second * 20)
}

type SyncMap struct {
	m  map[string]string
	l  sync.Mutex
	cm map[string]chan struct{}
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		m:  make(map[string]string),
		cm: make(map[string]chan struct{}),
	}
}
func (s *SyncMap) put(key, value string) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = value
	mc, ok := s.cm[key]
	if !ok {
		return
	}
	//获取执行面向对象的思想，自己定义个chan，里面的close方法用sync.once去修饰
	_, notClose := <-mc
	if notClose {
		close(s.cm[key])
	}
}
func (s *SyncMap) Get(key string) (string, error) {
	s.l.Lock()
	value, ok := s.m[key]
	if ok {
		s.l.Unlock()
		return value, nil
	}
	ch, ok := s.cm[key]
	if !ok {
		ch = make(chan struct{})
		s.cm[key] = ch
	}
	s.l.Unlock()
	tCtx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()
	select {
	case <-tCtx.Done():
		return "", tCtx.Err()
	case <-ch:
		break
	}
	s.l.Lock()
	defer s.l.Unlock()
	return s.m[key], nil
}
