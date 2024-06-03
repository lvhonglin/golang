package test

import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	var sMp sync.Map
	sMp.Store("key1", "val1")
	sMp.Store("key2", "val2")
	sMp.Store("key3", "val3")
	value, _ := sMp.Load("key1")
	println(value)
	a := 0
	sMp.Range(func(key, value any) bool {
		a++
		fmt.Printf("key=%v,val=%v\n", key, value)
		if a == 3 {
			return false
		}
		return true
	})
}

func Test2(t *testing.T) {
	var m sync.Map
	m.Store("123", nil)
	value, ok := m.Load("123")
	println(value)
	println(ok)
}
