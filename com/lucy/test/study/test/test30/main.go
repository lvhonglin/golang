package main

func main() {
	b := wordBreak("leetcode", []string{"leet", "code"})
	println(b)
}
func wordBreak(s string, wordDict []string) bool {
	chars := []byte(s)
	dp := make([]bool, len(chars))
	mmap := make(map[string]struct{}, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		mmap[wordDict[i]] = struct{}{}
	}
	for i := 0; i < len(chars); i++ {
		if _, ok := mmap[s[0:i+1]]; ok {
			dp[i] = true
		} else {
			for j := 0; j < i; j++ {
				if dp[j] == true {
					if _, ok := mmap[s[j+1:i+1]]; ok {
						dp[i] = true
					}
				}
			}
		}
	}
	return dp[len(chars)]
}
