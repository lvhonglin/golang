package main

import (
	"fmt"
	"strings"
)

func main() {

	join := strings.Join([]string{"1", "2"}, ".")
	fmt.Println(join)
}
