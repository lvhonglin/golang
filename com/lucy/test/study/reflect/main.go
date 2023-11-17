package main

import (
	model1 "golang/com/lucy/test/study/reflect/model"
	"reflect"
)

func main() {
	//获取结构体的字段的tags
	user := model1.GetUser()
	of := reflect.TypeOf(&user)
	field, exist := of.Elem().FieldByName("name")
	if exist {
		tag := field.Tag.Get("yaml")
		println(tag)
	}
}
