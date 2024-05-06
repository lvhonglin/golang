package main

import "strings"

func main() {
	//AKIDeDajeQzLlDmEWjyESpCiIzmiyAFNar13
	//zrjRvGxjMsPqAIILK6ueCoXBh9VvGfzP

	//DajeQzLlDmEWjyESpCiIzmiyAFNar13AKIDe
	//xjMsPqAIILK6ueCoXBh9VvGfzPzrjRvG

	println(computeStr("DajeQzLlDmEWjyESpCiIzmiyAFNar13AKIDe", 5))
	println(computeStr("xjMsPqAIILK6ueCoXBh9VvGfzPzrjRvG", 6))
}
func computeStr(s string, n int) string {
	builder := strings.Builder{}
	builder.WriteString(s[len(s)-n:])
	builder.WriteString(s[:len(s)-n])
	return builder.String()
}
func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i][0] = nums[i]
			dp[i][1] = nums[i]
		} else {
			dp[i][0] = getMin(nums[i], getMin(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
			dp[i][1] = getMax(nums[i], getMax(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		}
	}
	return getMax(0, getMax(dp[len(nums)-1][0], dp[len(nums)-1][1]))
}
func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
