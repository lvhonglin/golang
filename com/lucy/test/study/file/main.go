package main

import (
	"os"
	"path/filepath"
)

func main() {
	fileDir := filepath.Join("F:\\工作\\代码\\java\\testone\\guitar\\ideplugin\\perf_dataconsumer_外网", "src2/main")
	// 使用os.Stat来获取文件或目录的信息
	_, err := os.Stat(fileDir)

	// 判断目录是否存在
	if os.IsNotExist(err) {
		println("文件不存在")
	} else {
		println("文件存在")
	}
}
