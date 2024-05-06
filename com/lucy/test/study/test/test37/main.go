package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

func main() {
	sf := NewSafeSlice()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(a int) {
			sf.Add(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	println(len(sf.slice()))
}

type SafeSlice struct {
	p unsafe.Pointer
}

func NewSafeSlice() *SafeSlice {
	ms := make([]int, 0)
	return &SafeSlice{
		p: unsafe.Pointer(&ms),
	}
}

func (s *SafeSlice) Add(val int) {
	for {
		bak := s.p
		src := s.slice()
		new := make([]int, len(src)+1)
		copy(new, src)
		new[len(src)] = val
		newP := unsafe.Pointer(&new)
		swapped := atomic.CompareAndSwapPointer(&s.p, bak, newP)
		if swapped {
			break
		}
	}
}

func (s *SafeSlice) slice() []int {
	return *((*[]int)(s.p))
}
