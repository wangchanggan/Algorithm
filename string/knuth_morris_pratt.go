package string

func KnuthMorrisPratt(s, subStr string) bool {
	n := len(s)
	m := len(subStr)
	if n < m || (n == 0 && m > 0) {
		return false
	}
	if m == 0 {
		return true
	}

	next := getNexts(subStr, m)
	j := 0
	var i int
	for i = 0; i < len(s); i++ {
		for j > 0 && s[i] != subStr[j] {
			j = next[j-1] + 1
		}
		if s[i] == subStr[j] {
			j++
		}
		if j == m {
			return true
			//return i - m + 1
		}
	}
	return false
}

func getNexts(subStr string, subStrLength int) []int {
	next := make([]int, subStrLength)
	next[0] = -1
	k := -1
	var i int
	for i = 1; i < subStrLength; i++ {
		for k != -1 && subStr[k+1] != subStr[i] {
			k = next[k]
		}
		if subStr[k+1] == subStr[i] {
			k++
		}
		next[i] = k
	}
	return next
}
