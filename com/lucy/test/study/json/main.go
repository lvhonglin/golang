package main

import (
	"encoding/json"
)

type Username struct {
	Username string `json:"username"`
}
type User struct {
	Test string    `json:"test"`
	User *Username `json:"user"`
}

func main() {
	var d map[string]interface{}
	json.Unmarshal([]byte("{}"), &d)
	d["email"] = "ddd"
}
