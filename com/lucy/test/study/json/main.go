package main

import (
	"encoding/json"
	"io/ioutil"
)

type Username struct {
	Username string `json:"username"`
}
type User struct {
	Test string    `json:"test"`
	User *Username `json:"user"`
}

func main() {
	var uses []*User
	str, err := ioutil.ReadFile("C:\\Users\\Administrator\\GolandProjects\\golang\\com\\lucy\\test\\study\\json\\str.json")
	err = json.Unmarshal([]byte(str), &uses)
	if err != nil {
		println(err.Error())
		return
	}
	println(uses)
}
