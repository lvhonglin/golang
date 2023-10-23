package model

type User struct {
	name string
}

func GetUser() User {
	return User{name: "hello"}
}

type User2 struct {
	name string
}

func GetUser2() *User2 {
	return &User2{name: "hello"}
}
