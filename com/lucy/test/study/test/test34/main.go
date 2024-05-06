package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	var value unsafe.Pointer
	old := unsafe.Pointer(nil)
	tmp := new(int)
	new := unsafe.Pointer(tmp)
	fmt.Println(old)
	fmt.Println(new)
	fmt.Printf("%p\n", tmp)
	// 尝试将 value 的值从 nil 替换为 new(int)
	swapped := atomic.CompareAndSwapPointer(&value, old, new)
	fmt.Println("Swapped:", swapped) // 输出: true
	fmt.Println("====")
	fmt.Printf("%p\n", (*int)(value))
	// 尝试将 value 的值从 new(int) 替换为 new(int)
	swapped = atomic.CompareAndSwapPointer(&value, new, old)
	fmt.Println("Swapped:", swapped) // 输出: false
}
