package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"


示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"


示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"


提示：
1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300]
*/

func decodeString(s string) string {

	if len(s) <= 0 {
		return s
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			//fmt.Println(string(stack))
			//出栈,并且将当前的字符和数字循环n次，在入栈
			tmp := make([]byte, 0)
			for len(stack) > 0 && unicode.IsLetter(rune(stack[len(stack)-1])) {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1] //将[吐出来
			}
			//反转
			for a, b := 0, len(tmp)-1; a < b; a, b = a+1, b-1 {
				tmp[a], tmp[b] = tmp[b], tmp[a]
			}

			//次数
			tmpTimes := make([]byte, 0) //循环多次找到次数
			for len(stack) > 0 && unicode.IsDigit(rune(stack[len(stack)-1])) {
				tmpTimes = append(tmpTimes, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1] //将[吐出来
			}

			for a, b := 0, len(tmpTimes)-1; a < b; a, b = a+1, b-1 {
				tmpTimes[a], tmpTimes[b] = tmpTimes[b], tmpTimes[a]
			}

			times, _ := strconv.Atoi(string(tmpTimes))

			//push tmp 入栈
			for k := 0; k < times; k++ {
				stack = append(stack, tmp...)
			}
			continue

		} else if unicode.IsDigit(rune(s[i])) {
			stack = append(stack, s[i])
			continue
		} else if unicode.IsLetter(rune(s[i])) {
			stack = append(stack, s[i])
			continue
		} else if s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
	}

	return string(stack)
}

func main() {

	s := "2[abc]3[cd]ef"
	//s = "3[a2[c]]"
	//s = "100[leetcode]"
	//s = "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	fmt.Println(decodeString(s))

}
