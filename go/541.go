package main

import "fmt"

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


示例 1：
输入：s = "abcdefg", k = 2
输出："bacdfeg"

示例 2：
输入：s = "abcd", k = 2
输出："bacd"


提示：
1 <= s.length <= 104
s 仅由小写英文组成
1 <= k <= 104
*/

func reverseStr(s string, k int) string {

	if k == 0 || len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	for i := 0; i < len(runes); i += 2*k - 1 {
		end := i + k - 1
		if end >= len(runes) {
			end = len(runes) - 1
		}
		//交换
		//fmt.Println(i, end, string(runes))
		for i < end {
			tmp := runes[i]
			runes[i] = runes[end]
			runes[end] = tmp
			i++
			end--
		}
		//fmt.Println(i, end, string(runes))
	}

	return string(runes)
}

func main() {
	str1 := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	//str1:="abcdefg"
	//fmt.Println(reverseStr(str1, 2))
	res := reverseStr(str1, 39)
	result := "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi"
	fmt.Println(res == result)
}
