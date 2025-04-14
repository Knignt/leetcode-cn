package main

import "fmt"

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	i, j := 0, 0
	result := 0
	hashMap := make(map[byte]int, 0)

	for ; j < len(s); j++ {
		if _, exists := hashMap[s[j]]; exists {

			if j-i > result {
				result = j - i
			}

			for hashMap[s[i]] != 0 {
				delete(hashMap, s[i])
				i++
			}

		}
		hashMap[s[j]]++
	}
	if result == 0 && len(hashMap) > 0 {
		result = len(hashMap)
	}

	if j-i > result {
		result = j - i
	}

	return result
}

func main() {
	s := "aab"   //3
	s = "pwwkew" //3
	s = "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
