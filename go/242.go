package main

import (
	"fmt"
	"strings"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的
字母异位词
。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
*/

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) != len(t) {
		return false
	}

	sss := strings.Split(s, "")
	ttt := strings.Split(t, "")

	s1 := make(map[string]int)
	for i := 0; i < len(sss); i++ {
		s1[sss[i]]++
	}
	for j := 0; j < len(ttt); j++ {
		if s1[ttt[j]] == 0 {
			return false
		}
		s1[ttt[j]]--
	}

	return true
}

func main242() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
