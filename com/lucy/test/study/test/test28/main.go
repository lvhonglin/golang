package main

func main() {
	i := partition("bb")
	println(len(i))
}
func partition(s string) [][]string {
	chars := []byte(s)
	if len(chars) == 0 {
		return [][]string{{s}}
	}
	dp := make([][]bool, len(chars))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(dp))
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp)-i; j++ {
			if i < 3 {
				if chars[j] == chars[j+i] {
					dp[j][j+i] = true
				}
			} else {
				if dp[j+1][j+i-1] && chars[j] == chars[j+i] {
					dp[j][j+i] = true
				}
			}
		}
	}
	res := make([][]string, 0)
	tmp := make([]string, 0)

	dfs(s, chars, dp, &res, &tmp, -1, -1)
	return res
}
func dfs(s string, chars []byte, dp [][]bool, res *[][]string, tmpRes *[]string, left, right int) {
	if left > len(chars)-1 || right > len(chars)-1 || left > right {
		return
	}
	if (left == -1 && right == -1) || dp[left][right] {
		if right == len(chars)-1 {
			tmp := make([]string, len(*tmpRes))
			copy(tmp, *tmpRes)
			*res = append(*res, tmp)
			return
		}
		for i := 1; i < len(chars); i++ {
			if right+i > len(chars)-1 {
				break
			}
			*tmpRes = append(*tmpRes, s[right+1:right+i+1])
			dfs(s, chars, dp, res, tmpRes, right+1, right+i)
			*tmpRes = (*tmpRes)[0 : len(*tmpRes)-1]
		}
	}
}
