package main

import (
	"fmt"
	"os"
)

func main() {
	environ := os.Environ()
	for _, value := range environ {
		fmt.Println(value)
	}
}
