package longestsubstring

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(str string) int {
	var l = 0
	var r = -1
	var res int
	var freq = make([]int, 256)

	// 整个循环从 l == 0; r == -1 这个空窗口开始
	// 到l == s.size(); r == s.size()-1 这个空窗口截止
	for l < len(str) {
		if r+1 < len(str) && freq[str[r+1]] == 0 {
			r++
			freq[str[r]]++
		} else { // r = len(str)-1 || freq[str[r+1]] == 1
			freq[str[l]]--
			l++
		}

		res = max(res, r-l+1)
	}
	return res
}
