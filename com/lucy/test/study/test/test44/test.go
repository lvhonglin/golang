package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		v, ok := m[prerequisites[i][0]]
		if ok {
			m[prerequisites[i][0]] = append(v, prerequisites[i][1])
		} else {
			s := make([]int, 0)
			s = append(s, prerequisites[i][1])
			m[prerequisites[i][0]] = s
		}

	}
	resM := make(map[int]struct{})
	for i := 0; i <= numCourses-1; i++ {
		dfs(numCourses, make(map[int]struct{}), resM, i, m)
	}
	if len(resM) >= numCourses {
		return true
	} else {
		return false
	}
}

func dfs(numCourses int, set map[int]struct{}, resM map[int]struct{}, key int, m map[int][]int) bool {
	if _, ok := resM[key]; ok {
		return true
	}
	if len(resM) >= numCourses {
		return true
	}
	if _, ok := set[key]; ok {
		return false
	}
	set[key] = struct{}{}
	v, ok := m[key]
	flag := true
	if !ok {
		resM[key] = struct{}{}
	} else {
		for i := 0; i < len(v); i++ {
			res := dfs(numCourses, set, resM, v[i], m)
			if res == false {
				flag = false
				break
			}
		}
		if flag {
			resM[key] = struct{}{}
		}
	}
	delete(set, key)
	return flag
}
