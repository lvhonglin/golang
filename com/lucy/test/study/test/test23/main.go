package main

type A struct {
	arr [2]int
}

func longestConsecutive(nums []int) int {
	mmap := make(map[int]struct{})
	for _, v := range nums {
		mmap[v] = struct{}{}
	}
	max := 0
	for _, v := range nums {
		if _, ok := mmap[v+1]; !ok {
			cur := 1
			for {
				v = v - 1
				if _, ok := mmap[v]; ok {
					cur++
				} else {
					break
				}
			}
			max = getMax(max, cur)
		}
	}
	return max
}
func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
