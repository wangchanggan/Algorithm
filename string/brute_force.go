package string

//暴力匹配算法，也叫朴素匹配算法
func BruteForce(s, subStr string) bool {
	n := len(s)
	m := len(subStr)
	if n < m || (n == 0 && m > 0) {
		return false
	}
	if m == 0 {
		return true
	}

	for i := 0; i <= n-m; i++ {
		k := i
		for j := 0; j < m; j++ {
			if subStr[j] != s[k] {
				break
			}

			k++

			if j == m-1 {
				return true
			}
		}
	}

	return false
}
