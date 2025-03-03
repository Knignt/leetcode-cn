package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true



提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

/*
一个hash table 里面存储的是 括号的映射
stack 定义的是栈
如果是左括号,直接进入栈，有括号的话，查看当前的栈顶的左括号对应的是否和当前的括号一致即可，如果一致则将栈顶的数据弹出。
*/

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 { //如果是奇数项,必然不是ok
		return false
	}

	mapTable := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	stack := make([]byte, 0)

	i := 0
	for i < len(s) {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && s[i] == mapTable[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		i++
	}

	return len(stack) == 0
}

func main20() {
	s := "([])"
	fmt.Println(isValid(s))

}
