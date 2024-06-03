package test1

import (
	"fmt"
	"testing"
)

// 问：判断如下两个切片的地址是否相同？
// 答：结果不同，因为golang里面都是值传递，所以通过函数参数传递，
// 会创建一个新的指针
func TestA(t *testing.T) {
	a := append([]int{}, 1)
	println(&a)
	test(a)
}
func test(s []int) {
	println(&s)
}

func TestB(t *testing.T) {
	a := []int{}
	//len=0,cap=0
	fmt.Printf("1.len=%v,cap=%v\n", len(a), cap(a))
	//len=1,cap=1
	a = append(a, 1)
	fmt.Printf("2.len=%v,cap=%v\n", len(a), cap(a))
	//len=2,cap=2
	a = append(a, 1)
	fmt.Printf("3.len=%v,cap=%v\n", len(a), cap(a))
	//len=3,cap=4
	a = append(a, 1)
	fmt.Printf("4.len=%v,cap=%v\n", len(a), cap(a))
	//len=3,cap=3
	b := []int{1, 1, 1}
	fmt.Printf("5.len=%v,cap=%v\n", len(b), cap(b))
	//len=3,cap=3
	c := make([]int, 3)
	fmt.Printf("6.len=%v,cap=%v\n", len(c), cap(c))
	//len=4,cap=6
	c = append(c, 1)
	fmt.Printf("7.len=%v,cap=%v\n", len(c), cap(c))
	//len=3,cap=3
	d := make([]int, 0)
	e := []int{1, 2, 3}
	d = append(d, e...)
	fmt.Printf("8.len=%v,cap=%v\n", len(d), cap(d))
	//len=257,cap=512
	f := make([]int, 256)
	f = append(f, 1)
	fmt.Printf("9.len=%v,cap=%v\n", len(f), cap(f))

	g := make([]int, 512)
	g = append(g, 1)
	fmt.Printf("10.len=%v,cap=%v\n", len(g), cap(g))

}
