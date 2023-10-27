package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("x 的类型是 %T\n", err)
			println(err)
		}
	}()
	src := make(map[string]string)
	src["username"] = "test"
	src["password"] = "admin"
	user := User{}
	metadata := &mapstructure.Metadata{}
	config := &mapstructure.DecoderConfig{
		Result:           &user,
		WeaklyTypedInput: true,
		TagName:          "yaml",
		Metadata:         metadata,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
		return
	}
	err = decoder.Decode(&src)
	if err != nil {
		panic(err)
		return
	}
	marshal, _ := json.Marshal(user)
	println(string(marshal))
}
