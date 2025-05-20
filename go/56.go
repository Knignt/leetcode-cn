package main

import (
	"fmt"
	"slices"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。


提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })

	i := 1
	ans = append(ans, intervals[0])
	for i < len(intervals) {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1][1] = max(intervals[i][1], ans[len(ans)-1][1])
		} else {
			ans = append(ans, intervals[i])
		}
		i++
	}

	return ans
}

func main() {

	a := [][]int{
		{1, 3}, {8, 10}, {15, 18}, {2, 6},
	}

	a = [][]int{{1, 4}, {4, 5}}

	fmt.Println(merge(a))

}
