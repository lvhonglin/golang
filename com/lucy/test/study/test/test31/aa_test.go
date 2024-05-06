package main

import "testing"

func main() {

}
func Test_slice(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice2(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func Test_slice3(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

// 对s1修改会影响s
func Test_slice4(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	t.Logf("s: %v", s)
}

func Test_slice5(t *testing.T) {
	s := make([]int, 10, 12)
	v := s[10]
	// 求问，此时数组访问是否会越界
	t.Logf("%v", v)
}

// s1的修改会影响s，但是对s1的新增不会影响s
func Test_slice6(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	s1[0] = -1
	//s1是的len=1，cap=4，[-1]
	//s1会扩容：
	s1 = append(s1, []int{10, 11, 12, 13}...)
	//v := s[10]
	// ...
	// 求问，此时数组访问是否会越界
	//t.Logf("%v", v)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
func Test_slice7(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice(s1)
	t.Logf("s: %v", s)
}

func changeSlice(s1 []int) {
	s1[0] = -1
}
func Test_slice8(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice8(s1)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice8(s1 []int) {
	s1 = append(s1, 10)
}

// 百思不得其解
func Test_slice9(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	a := s
	//s[:2]的len=2，cap=2。s变成了：[0,1,3,4,4]
	s = append(s[:2], s[3:]...)
	t.Logf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	//v := s[4]
	// 是否会数组访问越界
	t.Logf("s: %v", s)
	t.Logf("s: %v", a)
}
func Test_slice10(t *testing.T) {
	s := make([]int, 512)
	//256——>512，
	s = append(s, 1)
	t.Logf("len of s: %d, cap of s: %d", len(s), cap(s))
}
func Test_slice11(t *testing.T) {
	s := make([]int, 3, 12)
	t.Logf("len of s: %d, cap of s: %d", len(s), cap(s))
}
func Test_slice12(t *testing.T) {
	s := make([]int, 10)
	s1 := s[4:]
	t.Logf("len of s: %d, cap of s: %d", len(s), cap(s))
	t.Logf("len of s: %d, cap of s: %d", len(s1), cap(s1))
}
