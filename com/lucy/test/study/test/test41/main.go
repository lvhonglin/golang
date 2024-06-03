package main

func main() {
	number := missingNumber([]int{3, 0, 1})
	println(number)
}
func missingNumber(nums []int) int {
	end := false
	res := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == len(nums) {
			end = true
			nums[i] = -1
			res = i
		} else {
			if nums[i] == i || nums[i] == -1 {
				continue
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
				if nums[i] == -1 {
					res = i
				}
				i--

			}
		}
	}
	if !end {
		return len(nums)
	}
	return res
}
