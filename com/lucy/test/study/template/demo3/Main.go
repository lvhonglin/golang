package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "test", Age: 123}
	//1.with的用法是设置他内部的{{.}}其中.的值
	tmpl, _ := template.New("test").Parse("{{with `人才`}}{{.}}{{end}}完事了")
	tmpl.Execute(os.Stdout, user)

}
