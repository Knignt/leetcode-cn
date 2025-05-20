package main

import (
	"fmt"
)

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/

func findAnagrams(s string, p string) []int {
	if len(s) <= 0 || len(p) <= 0 || len(p) > len(s) {
		return nil
	}

	ans := make([]int, 0)

	hashMap := make(map[byte]struct{}, 0)
	store := make(map[byte]struct{}, 0)
	for i := 0; i < len(p); i++ {
		hashMap[p[i]] = struct{}{}
		store[p[i]] = struct{}{}
	}

	i, j := 0, 0
	for j < len(s) {
		if j-i < len(p) {
			delete(hashMap, s[j])
			j++
			continue
		}

		if len(hashMap) == 0 {
			ans = append(ans, i)
		}
		i++
		j = i
		hashMap = store
	}

	if j-i >= len(s) {
		hashMap = store
		last := i

		for i < j {
			delete(hashMap, s[i])
			i++
		}

		if len(hashMap) == 0 {
			ans = append(ans, last)
		}
	}

	return ans
}

func main() {
	s := "cbaebabacd"
	s = "abab"
	a := "abc"
	a = "ab"
	fmt.Println(findAnagrams(s, a))
}
