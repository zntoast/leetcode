package lcp

//01.猜数字
func Game(guess []int, answer []int) int {
	count := 0
	for j := 0; j < len(answer); j++ {
		if guess[j] == answer[j] {
			count++
		}
	}
	return count
}
