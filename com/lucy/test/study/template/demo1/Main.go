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
	tmpl, _ := template.New("test").Parse("{{.Name}} 的 age 是 {{.Age}}")
	tmpl.Execute(os.Stdout, user)
	println()
	tmpl, _ = template.New("test").Parse("xxx {{.}} xxx")
	tmpl.Execute(os.Stdout, user)
	println()
	//如果在{{}}中加入-，加在左面表示去掉左面的空格，加在右面表示去掉右面的空格
	tmpl, _ = template.New("test").Parse("xxx  {{- . -}}  xxx")
	tmpl.Execute(os.Stdout, user)
	println()

}
