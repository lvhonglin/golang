package main

func main() {
	a := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	trees := findMinHeightTrees(6, a)
	println(trees)
}
func findMinHeightTrees(n int, edges [][]int) []int {
	m := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		v, ok := m[edges[i][0]]
		if ok {
			v = append(v, edges[i][1])
			m[edges[i][0]] = v
		} else {
			v = make([]int, 1)
			v[0] = edges[i][1]
			m[edges[i][0]] = v
		}
		v, ok = m[edges[i][1]]
		if ok {
			v = append(v, edges[i][0])
			m[edges[i][1]] = v
		} else {
			v = make([]int, 1)
			v[0] = edges[i][0]
			m[edges[i][1]] = v
		}
	}

	res := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		dfs(m, i, i, make(map[int]struct{}), make(map[int]struct{}), 1, res)
	}
	min := n + 1
	for _, v := range res {
		if v < min {
			min = v
		}
	}
	result := make([]int, 0)
	for k, v := range res {
		if v == min {
			result = append(result, k)
		}
	}
	return result

}
func dfs(m map[int][]int, head int, n int, cur map[int]struct{}, set map[int]struct{}, num int, res map[int]int) {

	if _, ok := set[n]; ok {
		return
	}
	v1, ok := res[head]
	if ok {
		if num > v1 {
			res[head] = num
		}
	} else {
		res[head] = num
	}
	set[n] = struct{}{}

	v, _ := m[n]
	for i := 0; i < len(v); i++ {
		cur[v[i]] = struct{}{}
	}
	for i := 0; i < len(v); i++ {
		dfs(m, head, v[i], cur, set, num+1, res)
	}

}
