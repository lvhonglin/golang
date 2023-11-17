package main

import "os"

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	println(getwd)
	//切换当前目录
	err = os.Chdir("C:\\Users\\Administrator\\GolandProjects\\golang\\com\\lucy\\test\\study")
	if err != nil {
		return
	}
	getwd, err = os.Getwd()
	if err != nil {
		return
	}
	println(getwd)
}
