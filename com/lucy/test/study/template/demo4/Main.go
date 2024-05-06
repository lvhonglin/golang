package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Map1 map[string]interface{}
}
type A interface {
	count(num int)
}
type a struct {
}

func (*a) count(num int) {

}
func NewA() A {
	return &a{}
}

type B struct {
	tmpA A
}

func returnB() *B {
	return &B{
		tmpA: NewA(),
	}
}

func main() {
	mymap := make(map[string]interface{})
	mymap["var1"] = "我不知道"
	user := User{Name: "test", Age: 123, Map1: mymap}
	tmpl, _ := template.New("test").Parse("{{.Name}} 的 age 是 {{.Age}} 的map中的var1是 {{ .Map1.var1 }}")
	tmpl.Execute(os.Stdout, user)

}
