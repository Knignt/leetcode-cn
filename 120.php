<?php

/**
120. 三角形最小路径和
中等
1.3K
相关企业
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。



示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
2
3 4
6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10


提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104


进阶：

你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？

 **/

class Solution {

    /**
     * @param Integer[][] $triangle
     * @return Integer
     */
    function minimumTotal($triangle) {
        if(empty($triangle)) return 0;
        if(count($triangle) == 1) return max($triangle[0]);

        $cnt = count($triangle);
        for($i=$cnt-2;$i>=0;$i--){
            for($j = 0;$j <= count($triangle[$i])-1;$j++){
                $triangle[$i][$j] = min($triangle[$i+1][$j], $triangle[$i+1][$j+1]) + $triangle[$i][$j];
            }
        }

        return $triangle[0][0];
    }
}
$a = new Solution();
$triangle = [[2],[3,4],[6,5,7],[4,1,8,3]];
var_dump($a->minimumTotal($triangle));


