package main

import "fmt"

/*
车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。

示例 1：
输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false

示例 2：
输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true

提示：
1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi <= 100
0 <= fromi < toi <= 1000
1 <= capacity <= 105
*/

/*
差分数组思想:
仅仅只需要关注节点开始以及节点结束，节点开始时 加上值， 节点结束时 减去值
为什么中间的值不需要管？
后续在遍历时，中间的值为0，表明中间这里的话，没有增加人数，所以可以直接加
而对于下车时，这里是负数，会减去上车的人数；如果中间又有人上车，也仅仅只会减去刚开始上车的人数。
*/

func carPooling(trips [][]int, capacity int) bool {
	if capacity <= 0 || len(trips) <= 0 {
		return false
	}

	ltrips := len(trips)
	ansMap := make(map[int]int)
	for i := 0; i < ltrips; i++ {
		if _, exists := ansMap[trips[i][1]]; !exists {
			ansMap[trips[i][1]] = 0
		}
		ansMap[trips[i][1]] += trips[i][0]

		if _, exists := ansMap[trips[i][2]]; !exists {
			ansMap[trips[i][2]] = 0
		}
		ansMap[trips[i][2]] -= trips[i][0]
	}

	cur := 0
	for i := 0; i < 1001; i++ {
		if _, exist := ansMap[i]; exist {
			cur += ansMap[i]
		}

		if cur > capacity {
			return false
		}
	}

	return true
}

func main() {

	a := [][]int{
		{2, 1, 5},
		{3, 3, 7},
	}

	fmt.Println(carPooling(a, 4))

}
