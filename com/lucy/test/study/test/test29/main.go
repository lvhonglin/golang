package main

func main() {
	circuit := canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
	println(circuit)
}
func canCompleteCircuit(gas []int, cost []int) int {
	gasCount := 0
	for i := 0; i < len(gas); i++ {
		gasCount += gas[i]
	}
	costCount := 0
	for i := 0; i < len(cost); i++ {
		costCount += cost[i]
	}
	if gasCount < costCount {
		return -1
	}
	sum := 0
	init := false
	maxIndex := -1
	for i := 0; i < len(gas); i++ {
		sum += (gas[i] - cost[i])
		if sum >= 0 {
			if init == false {
				init = true
				maxIndex = i
			}
		} else {
			sum = 0
			init = false
			maxIndex = -1
		}
	}
	return maxIndex
}
