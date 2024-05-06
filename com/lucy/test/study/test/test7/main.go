package main

import "errors"

var NOT_FOUND_ERR error = errors.New("not found")

func main() {
	err := returnErr()
	println(err == &NOT_FOUND_ERR)
}

func returnErr() *error {
	return &NOT_FOUND_ERR
}
