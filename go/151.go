package main

import (
	"fmt"
	"strings"
)

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：
输入：s = "the sky is blue"
输出："blue is sky the"

示例 2：
输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。

示例 3：
输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词

进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
*/

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	return " "
}

func reverseWords111(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes = append([]rune{32}, runes...)
	//从后往前找
	var result []rune //结果定义
	j := len(runes) - 1
	for runes[j] == 0 { //从后往前找,找到第一个不为空的字符
		j--
	}
	right := j

	for j >= 0 {
		//找到第一个字
		if runes[j] != 32 {
			j--
			continue
		} else {

			//将j 到right之间的字全部放到result里面去
			if len(result) == 0 {
				result = append(runes[j : right+1])
			} else {
				result = append(result, runes[j:right+1]...)
			}

			//往前找到第一个字,并且挪动right
			for j > 0 && runes[j] == 32 {
				j--
			}
			right = j
			j--
		}
	}

	return strings.Trim(string(result), " ")
}

func main() {
	fmt.Println(reverseWords("the sky is good"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords(" asdasd df f"))
}
