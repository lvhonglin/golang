package main

import "encoding/json"

func main() {
	a := []string{"aa", "bb"}
	marshal, err := json.Marshal(&a)
	if err != nil {
		return
	}
	str := string(marshal)
	println(str)
	b := make([]string, 0)
	err = json.Unmarshal([]byte(str), &b)
	if err != nil {
		return
	}
	println(b[0])

}
