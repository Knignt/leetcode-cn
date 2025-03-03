<?php

/**

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 200
 **/

class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function minPathSum($grid) {
        if(empty($grid)) return 0;

        $dp = [];
        $dp[0][0] = $grid[0][0];
        for($i=0;$i<count($grid);$i++){
            for($j=0;$j<count($grid[$i]);$j++){
                if(!isset($dp[$i][$j])) $dp[$i][$j] = 0;

                if($i == 0){
                    if($j > 0){
                        $dp[$i][$j] = $dp[$i][$j-1] + $grid[$i][$j];
                    }
                }
                if($j == 0){
                    if($i > 0){
                        $dp[$i][$j] = $dp[$i-1][$j] + $grid[$i][$j];
                    }
                }
            }
        }

        for($i=1;$i<count($grid);$i++){
            for($j=1;$j<count($grid[$i]);$j++){
                $dp[$i][$j] = min($dp[$i-1][$j], $dp[$i][$j-1])+$grid[$i][$j];
            }
        }


        return $dp[$i-1][$j-1];
    }
}

$a = new Solution();
$grid = [[1,2,3],[4,5,6]];
var_dump($a->minPathSum($grid));

