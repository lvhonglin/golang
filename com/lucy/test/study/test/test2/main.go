package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 16.32999992370605342342342342345
	str := fmt.Sprintf("%.2f", f)
	ff, err := strconv.ParseFloat(str, 64)
	if err != nil {
	}
	println(float32(ff))
}
