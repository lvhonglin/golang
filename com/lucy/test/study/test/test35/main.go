package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type SafeSlice struct {
	data unsafe.Pointer
}

func NewSafeSlice(cap int) *SafeSlice {
	return &SafeSlice{
		data: unsafe.Pointer(&[]interface{}{}),
	}
}
func (s *SafeSlice) Append(item interface{}) {
	for {
		oldPointer := s.data
		old := s.Load()
		newSlice := make([]interface{}, len(old)+1)
		copy(newSlice, old)
		newSlice[len(old)] = item
		if atomic.CompareAndSwapPointer(&s.data, oldPointer, unsafe.Pointer(&newSlice)) {
			break
		}
	}
}
func (s *SafeSlice) Load() []interface{} {
	//用atomic.LoadPointer是为了保证s.data的可见性
	return *(*[]interface{})(atomic.LoadPointer(&s.data))
}
func main() {
	safeSlice := NewSafeSlice(0)
	safeSlice.Append(1)
	safeSlice.Append(2)
	safeSlice.Append(3)

	fmt.Println(safeSlice.Load()) // [1 2 3]
}
