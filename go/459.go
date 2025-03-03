package main

import (
	"fmt"
	"strings"
)

/*
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

示例 1:
输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。

示例 2:
输入: s = "aba"
输出: false

示例 3:
输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)

提示：
1 <= s.length <= 104
s 由小写英文字母组成
*/

/*
判断字符串s是否由重复子串组成，只要两个s拼接在一起，里面还出现一个s的话，就说明是由重复子串组成。
如果不是重复的串，那么在下标从1开始到len(s+s)-1的位置就不会包含s
如果是重复的串，那么下标从1开始到len(s+s)-1就会出现包含的s
*/

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}

func main() {
	s := "abaababaab"
	fmt.Println(repeatedSubstringPattern(s))
}
