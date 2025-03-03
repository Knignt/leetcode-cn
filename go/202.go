package main

import "fmt"

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
示例 2：
输入：n = 2
输出：false

提示：

1 <= n <= 231 - 1
*/

func getSum(n int) int {
	if n == 0 {
		return 0
	}

	if n < 10 {
		return n * n
	}

	result := 0
	for n > 0 {
		tmp := n % 10
		result += tmp * tmp
		n /= 10
	}

	return result
}
func isHappy(n int) bool {
	if n == 1 || n == 10 || n == 100 {
		return true
	}
	var map1 = make(map[int]struct{})
	//判断某一个结果是否出现过,如果出现过，那么就表示进入了循环,否则就一直找下去
	//为什么如果出现过就会是进入循环,原因在于比如计算出来32，那么下次还是3^2+2^2求和，
	//如果有相等，以后还是这么计算，这样就是重复的意思
	for true {
		tmp := getSum(n)
		fmt.Println(tmp)
		if tmp == 1 {
			return true
		}
		if _, exists := map1[tmp]; exists {
			return false
		}
		map1[tmp] = struct{}{}
		n = tmp
	}

	return true
}

func main202() {
	fmt.Println(isHappy(2))
}
