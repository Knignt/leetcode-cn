package main

/*
给出由小写字母组成的字符串 s，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 s 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。


示例：
输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。
之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。


提示：
1 <= s.length <= 105
s 仅由小写英文字母组成。
*/

/*
删除字符串中的所有相邻重复项

主要思路是：如果是一直相邻，类似abba这种的话，可以使用栈来处理,如果不一致，直接给栈里面放
如果是相邻的直接pop栈，最后栈里面剩余的就是 结果
*/

func removeDuplicates1(s string) string {
	if len(s) == 0 {
		return ""
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func main1047() {

}
