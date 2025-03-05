package main

import (
	"fmt"
)

/*
难度分:1263
给你字符串 s 和整数 k 。

请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

英文中的 元音字母 为（a, e, i, o, u）。



示例 1：

输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母。
示例 2：

输入：s = "aeiou", k = 2
输出：2
解释：任意长度为 2 的子字符串都包含 2 个元音字母。
示例 3：

输入：s = "leetcode", k = 3
输出：2
解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
示例 4：

输入：s = "rhythms", k = 4
输出：0
解释：字符串 s 中不含任何元音字母。
示例 5：

输入：s = "tryhard", k = 4
输出：1


提示：

1 <= s.length <= 10^5
s 由小写英文字母组成
1 <= k <= s.length

*/

/*
解法：
1、j从0开始到k计算出当前一共有多少的元音字符，为滑动窗口
2、i从0开始+，如果碰到元音字符，直接--，j也开始加，碰到直接--
3、在中间不停的对比，将最大值记下来并且返回
*/

// leetcode 时间复杂度来看，字符对比 比 哈希效率更快
func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func maxVowels(s string, k int) int {
	if len(s) == 0 || k <= 0 {
		return 0
	}

	i, j := 0, 0
	result := 0
	strLen := len(s)
	for j < k {
		if isVowel(s[j]) {
			result++
		}
		j++
	}
	if strLen == k {
		return result
	}

	maxResult := result
	for j < strLen {
		//
		if isVowel(s[i]) {
			result--
		}
		i++

		if isVowel(s[j]) {
			result++
		}
		j++

		if maxResult < result {
			maxResult = result
		}
		if result == k {
			return k
		}
	}

	return maxResult
}

func main1456() {

	fmt.Println(maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33))

}
