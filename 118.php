<?php

/**
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]


提示:

1 <= numRows <= 30
 **/

class Solution {

    /**
     * @param Integer $numRows
     * @return Integer[][]
     */
    function generate($numRows) {
        if($numRows == 0) return [];

        $dp = [];
        $dp[0] = [1];
        for($i=1;$i<$numRows;$i++){
            $dp[$i] = [];
            for($j=0;$j<$i+1;$j++){
                if(isset($dp[$i-1][$j-1]) && isset($dp[$i-1][$j])){
                    $dp[$i][$j] = $dp[$i-1][$j-1] + $dp[$i-1][$j];
                }else{
                    $dp[$i][$j] = 1;
                }
            }
        }

        return $dp;
    }
}

$a = new Solution();
$numRows = 5;
var_dump($a->generate($numRows));
