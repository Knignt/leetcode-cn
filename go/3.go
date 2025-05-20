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
		//这里如果之前存在过,那么找到相等的点,然后将 左边指针移除之前的滑动窗口，再右边指针加入。
		if _, exists := hashMap[s[j]]; exists {

			if j-i > result { //只要有相等，那么就直接求长度即可。
				result = j - i
			}

			for s[i] != s[j] { //找到在当前窗口内，相等的位置
				delete(hashMap, s[i])
				i++
			}

			//然后左边+1，并且将当前的值在hashMap里面减去
			delete(hashMap, s[i])
			i++
		}
		if j < len(s) {
			hashMap[s[j]]++
		}
	}

	//整个串都是非重复子串
	if result == 0 && len(hashMap) == len(s) {
		result = len(hashMap)
	}

	if j-i > result {
		result = j - i
	}

	return result
}

func main() {
	s := "aab"   //2
	s = "pwwkew" //3
	s = "dvdf"   //3
	fmt.Println(lengthOfLongestSubstring(s))
}
