package main

import (
	"fmt"
	"strings"
)

func main() {
	items := []string{"item1", "item2", "item3"}

	// 添加双引号包裹每个字符串元素
	var quotedItems []string
	for _, item := range items {
		quotedItems = append(quotedItems, fmt.Sprintf(`"%s"`, item))
	}

	// 将带有双引号的字符串切片转换成逗号分隔的字符串
	inClause := strings.Join(quotedItems, ",")
	res := fmt.Sprintf("(%s)", inClause)
	println(res)
}
