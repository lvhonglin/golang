package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()

	// 获取上一个月的时间
	lastMonth := now.AddDate(0, 0, -5)

	// 按指定格式格式化日期和时间
	lastMonthTime := lastMonth.Format("2006-01-02 15:04:05")
	println(lastMonthTime)
	layout := "2006-01-02 15:04:05"

	// 解析时间字符串为 time.Time 类型
	time1, _ := time.Parse(layout, lastMonthTime)
	time2, _ := time.Parse(layout, "2024-01-02  15:04:05")

	if time2.Before(time1) {
		fmt.Println("time2 is before time1")
	}

}
