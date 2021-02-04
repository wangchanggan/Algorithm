package string

import "math"

const SIZE = 256

//在模式串与主串匹配的过程中，当模式串和主串某个字符不匹配的时候，能够跳过一
//些肯定不会匹配的情况，将模式串往后多滑动几位。
func BoyerMoore(s, subStr string) bool {
	n := len(s)
	m := len(subStr)
	if n < m || (n == 0 && m > 0) {
		return false
	}
	if m == 0 {
		return true
	}

	bc := make([]int, SIZE)   //记录模式串中每个字符最后出现的位置
	generateBC(subStr, m, bc) //构建坏字符哈希表

	suffix := make([]int, m)
	prefix := make([]bool, m)
	generateGS(subStr, m, suffix, prefix)

	i := 0
	for i <= n-m {
		var j int
		//模式串从后往前匹配
		for j = m - 1; j >= 0; j-- {
			//坏字符串对应模式串中的小标是j
			if s[i+j] != subStr[j] {
				break
			}
		}
		if j < 0 {
			return true
			//return i     //匹配成功，返回主串与模式串第一个匹配的字符的位置
		}
		x := j - bc[int(s[i+j])]
		y := 0
		if j < m-1 {
			y = moveByGS(j, m, suffix, prefix)
		}
		i = i + int(math.Max(float64(x), float64(y)))
	}
	return false
}

func generateBC(subStr string, m int, bc []int) {
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}

	for i := 0; i < m; i++ {
		ascii := int(subStr[i])
		bc[ascii] = i
	}
}

func generateGS(subStr string, m int, suffix []int, prefix []bool) {
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ {
		j := i
		k := 0 //公共后缀子串长度
		for j >= 0 && subStr[j] == subStr[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
}

func moveByGS(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j //好后缀长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-2; r++ {
		if prefix[m-r] == true {
			return r
		}
	}
	return m
}
