package main

import "golang/com/lucy/test/study/flag/cobra/cmd"

func main() {
	//main方法在项目根目录
	//用go build -o hello，打包成hello
	//打包之后执行：.\hello.exe echo time helloworld da -t 2 -k true
	//打包之后执行：.\hello.exe echo time helloworld da -t 2 --kaiguan
	cmd.Execute()
}
