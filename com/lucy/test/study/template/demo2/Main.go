package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Age   int
	Citys []*City
}
type City struct {
	Name string
	Gdp  float64
}

func main() {
	city1 := &City{Name: "大连", Gdp: 222}
	city2 := &City{Name: "北京", Gdp: 444}
	user := User{Name: "test", Age: 123, Citys: []*City{city1, city2}}
	//1.在模板中定义内部变量
	tmpl, _ := template.New("test").Parse("" +
		"{{$var:=2150}}Internal veriable has value [{{$var}}]\n")
	tmpl.Execute(os.Stdout, user)
	println("=================")
	//2.if else，eq是等于，ne是不等于,lt是小于，le是小于等于,gt是大于，ge是大于等于
	tmpl, _ = template.New("test").Parse("" +
		"{{ if eq .Name `test`}}" +
		"This is filter..." +
		"{{else}}" +
		"it is something else..." +
		"{{end}}\n" +
		"完事了" +
		"\n")
	tmpl.Execute(os.Stdout, user)
	println("==================")
	//3.range遍历
	tmpl, _ = template.New("test").Parse("My city are:\n{{range .Citys}}{{.Name}}的gdp是{{.Gdp}}\n{{end}}")
	tmpl.Execute(os.Stdout, user)
	//4.index,element遍历
	tmpl, _ = template.New("test").Parse("My city are:\n" +
		"{{range $index1,$element1:= .Citys}}" +
		"index-{{$index1}}:{{$element1.Name}}的gdp是{{$element1.Gdp}}\n{{end}}")
	tmpl.Execute(os.Stdout, user)
	println("===================")
	//5.模板中使用外部的函数
	tmpl, _ = template.New("test").Funcs(template.FuncMap{"add1": add}).
		Parse("my city are:\n{{range $index1,$elem1:= .Citys}}" +
			"{{add1 $index1 1}}) {{$elem1.Name}}\n" +
			"{{end}}")
	tmpl.Execute(os.Stdout, user)
}
func add(a, b int) int {
	return a + b
}
