package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mymap := make(map[string]interface{})
	str := "{\"name\":\"hello1\",\"age\":\"123\"}"
	err := json.Unmarshal([]byte(str), &mymap)
	if err != nil {
		return
	}
	fmt.Println(mymap)
}
