package main

func main() {
	lis := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	println(lis)
}
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if max < dp[j] {
					max = dp[j]
				}
			}
		}
		dp[i] = max + 1
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
