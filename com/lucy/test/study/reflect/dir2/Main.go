package main

import (
	"fmt"
	model1 "golang/com/lucy/test/study/reflect/model"
	"reflect"
)

func main() {
	user := model1.GetUser()
	of := reflect.ValueOf(user)
	name := of.FieldByName("name")
	fmt.Println(name.String())

	user2 := model1.GetUser2()
	//获取指针类型必须加elem，不加的话会报错
	elem := reflect.ValueOf(user2).Elem()
	byName := elem.FieldByName("name")
	fmt.Println(byName.String())
}
