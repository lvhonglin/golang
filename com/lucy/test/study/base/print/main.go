package main

import (
	"encoding/json"
	"fmt"
)

type fileInfo struct {
	name    string
	content string
	baseDir string
}

var Guitar = fileInfo{
	name: "guitar.yml",
	content: `# 从 TestOne Web 快速获取：http://testone.woa.com/one-ci#/ci-base
product:
  product_uuid: xxx
  product_name: xxx
  business_uuid: xx
  business_name: xx

# 123 平台服务信息（建议配置）
server:
  app: "%s"
  server: "%s"%s
rick:
  public_protocol:
    paths: %s`,
}

func main() {
	arr := []string{"a/b", "b/c"}
	marshal, err := json.Marshal(&arr)
	if err != nil {
		return
	}
	content := fmt.Sprintf(Guitar.content, "hello", "world", "", string(marshal))
	println(content)
}
