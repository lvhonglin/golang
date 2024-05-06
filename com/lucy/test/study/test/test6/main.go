package main

import (
	"fmt"
)

type User struct {
	Parent
}
type Parent struct {
	ParentName string
	Age        int
}

func (p *Parent) Pay() {
	fmt.Println("pay")
}
func main() {
	var a *float32 = nil
	println(*a)
}
