package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	return ""
}

func main() {
	a := ""
	fmt.Println(longestPalindrome(a))
}
