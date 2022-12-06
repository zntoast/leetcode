package intermediate

/* 电话号码的字母组合 */
func LetterCombinations(digits string) []string {
	lenght := len(digits)
	// 枚举
	phoneMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	result := []string{}
	// 深度优先搜索
	dfs := func(digits string, index int, s string) {}
	dfs = func(digits string, index int, s string) {
		if lenght == index {
			result = append(result, s)
			return
		}
		// 获取号码
		digit := string(digits[index])
		// 获取号码对应的字母
		letters := phoneMap[digit]
		for _, v := range letters {
			dfs(digits, index+1, s+string(v))
		}
	}
	dfs(digits, 0, "")
	return result
}
