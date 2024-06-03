package main

func main() {
	topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v, ok := m[nums[i]]
		if ok {
			m[nums[i]] = v + 1
		} else {
			m[nums[i]] = 1
		}
	}
	mm := make(map[int][]int)
	for key, value := range m {

		v1, ok := mm[value]
		if ok {
			v1 = append(v1, key)
		} else {
			v1 = make([]int, 1)
			v1[0] = key
		}
		mm[value] = v1
	}
	res := make([]int, 0)
	index := 0
	for i := 10000; i >= 0; i-- {
		if k == 0 {
			break
		}
		v, ok := mm[i]
		if !ok {
			continue
		}
		res = append(res, v...)

		k -= len(v)

		index++
	}
	return res
}
