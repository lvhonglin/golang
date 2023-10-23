package main

import (
	"fmt"
	_ "golang/com/lucy/test/study/io/file/init"
	"gopkg.in/yaml.v3"
	ioutil "io/ioutil"
)

func main() {
	fileybytes, err := ioutil.ReadFile("com\\lucy\\test\\study\\io\\file\\filetobyte\\config.yml")
	if err != nil {
		return
	}
	config := &Config{}
	err = yaml.Unmarshal(fileybytes, config)
	if err != nil {
		return
	}
	fmt.Println(config.Name)
	fmt.Println(config.Version)
	fmt.Println(config.Age)
	fmt.Println(config.Address.Province)
}

type Config struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Version string   `yaml:"version"`
	Address *address `yaml:"address"`
}
type address struct {
	Province     string `yaml:"province"`
	Municipality string `yaml:"municipality"`
}
