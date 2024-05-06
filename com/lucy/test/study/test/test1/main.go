package main

import "strings"

func main() {
	s := " @alias=/rta/ams_sp\n"
	s = strings.TrimPrefix(s, " @alias=")
	suffix := strings.TrimSuffix(s, "\n")
	println(suffix)
}
