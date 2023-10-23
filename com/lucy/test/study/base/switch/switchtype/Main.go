package main

import "fmt"

type Persion interface {
	Name() string
}
type Student struct {
}

func (student *Student) Name() string {
	return "student"
}
func (student *Student) Study() {
	fmt.Println("study")
}

type Worker struct {
}

func (worker *Worker) Name() string {
	return "worker"
}
func (worker *Worker) Work() {
	fmt.Println("work")
}
func main() {
	var person Persion = &Student{}
	switch t := person.(type) {
	case *Student:
		fmt.Println("学生")
		t.Study()
	case *Worker:
		fmt.Println("工人")
		t.Work()
	default:
		fmt.Println("啥也不是")
	}
}
