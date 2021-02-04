package string

import (
	"crypto/md5"
)

//通过哈希算法对主串中的n-m+1个子串分别求哈希值，然后逐个与模式串的哈希值比较大小。
// 如果某个子串的哈希值与模式串相等，那就说明对应的子串和模式串匹配了。
func RabinKarp(s, subStr string) bool {
	n := len(s)
	m := len(subStr)
	if n < m || (n == 0 && m > 0) {
		return false
	}
	if m == 0 {
		return true
	}

	dstMd5 := MD5(subStr)
	for i := 0; i <= n-m; i++ {
		if MD5(s[i:i+m]) == dstMd5 {
			return true
		}
	}

	return false
}

//计算字符串MD5值
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
