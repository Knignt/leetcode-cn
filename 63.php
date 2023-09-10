<?php

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。



示例 1：


输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
示例 2：


输入：obstacleGrid = [[0,1],[0,0]]
输出：1


提示：

m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1

 **/

class Solution {

    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    function uniquePathsWithObstacles($obstacleGrid) {
        if(empty($obstacleGrid)) return 0;
        if($obstacleGrid[0][0] == 1) return 0 ;
        if($obstacleGrid[count($obstacleGrid)-1][count($obstacleGrid[0])-1] == 1) return 0;     //谁家好题目会给自己开始节点和结束节点设置障碍啊
        if($obstacleGrid[0][1] == 1 && $obstacleGrid[1][0] == 1) return 0;

        $dp = [];
        $dp[0][0] = 1;
        for($i=1;$i<count($obstacleGrid);$i++){
            if($obstacleGrid[$i][0] == 1){
                break;
            }else{
                $dp[$i][0] = 1;
            }
        }

        for($i=1;$i<count($obstacleGrid[0]);$i++){
            if($obstacleGrid[0][$i] == 1){
                break;
            }else{
                $dp[0][$i] = 1;
            }
        }

        for($i=1;$i<count($obstacleGrid);$i++){
            for($j=1;$j<count($obstacleGrid[$i]);$j++){
                if($obstacleGrid[$i][$j] == 1) {
                    $dp[$i][$j] = 0;
                }else{
                    $dp[$i][$j] = $dp[$i][$j-1] + $dp[$i-1][$j];
                }
            }
        }

        return $dp[$i-1][$j-1];
    }
}

$a = new Solution();
$obstacleGrid = [[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]];
var_dump($a->uniquePathsWithObstacles($obstacleGrid));


