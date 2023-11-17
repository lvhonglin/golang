package main

import "path/filepath"

func main() {
	a := "c:/data1/2"
	join := filepath.Join(a, "3")
	println(join)
}
