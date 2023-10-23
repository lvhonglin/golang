package main

import (
	"fmt"
	"strconv"
)

type myfunc func(int, int) int

func main() {
	var a int = 1
	b := strconv.Itoa(a)
	println("----")
	println(b)
	println("123123")
	fmt.Println(getSum(func(i int, i2 int) int {
		return i + i2
	}))
	fmt.Println(getSum(func(i int, i2 int) int {
		return i * i2
	}))

}
func getSum(f myfunc) int {
	return f(1, 2)
}
