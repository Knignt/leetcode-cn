<?php


/**
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。



示例 1：



输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
输出：13
解释：如图所示，为和最小的两条下降路径
示例 2：



输入：matrix = [[-19,57],[-40,-5]]
输出：-59
解释：如图所示，为和最小的下降路径


提示：

n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100
 **/

class Solution {

    /**
     * @param Integer[][] $matrix
     * @return Integer
     */
    function minFallingPathSum($matrix) {
        if(empty($matrix)) return 0;

        $dp = [];
        for($i=0;$i<count($matrix);$i++){
            for($j=0;$j<count($matrix[$i]);$j++){
                $dp[$i][$j] = 0;
                if($i == 0){
                    $dp[$i][$j] = $matrix[$i][$j];
                }
            }
        }

        for($i=1;$i<count($matrix);$i++){
            for($j=0;$j<count($matrix[$i]);$j++){
                if(!isset($dp[$i-1][$j]) ) $dp[$i-1][$j] = PHP_INT_MAX;
                if(!isset($dp[$i-1][$j+1])) $dp[$i-1][$j+1] = PHP_INT_MAX;
                if(!isset($dp[$i-1][$j-1])) $dp[$i-1][$j-1] = PHP_INT_MAX;
                $dp[$i][$j] = min($dp[$i-1][$j],$dp[$i-1][$j-1],$dp[$i-1][$j+1])+$matrix[$i][$j];
            }
        }

        $ret = PHP_INT_MAX;
        for($j=0;$j<count($matrix[$i-1]);$j++){
            $ret = min($ret, $dp[$i-1][$j]);
        }

        return $ret;
    }
}

$a = new Solution();
$matrix = [[2,1,3],[6,5,4],[7,8,9]];
var_dump($a->minFallingPathSum($matrix));

/**
题解：
 * 这道题目难点在于最后要在dp数组的最后一行获取最小值，因为题意是下降路径最小和
 **/
