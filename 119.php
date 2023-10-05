<?php
/**
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: rowIndex = 3
输出: [1,3,3,1]
示例 2:

输入: rowIndex = 0
输出: [1]
示例 3:

输入: rowIndex = 1
输出: [1,1]


提示:

0 <= rowIndex <= 33


进阶：

你可以优化你的算法到 O(rowIndex) 空间复杂度吗？

 **/

class Solution {

    /**
     * @param Integer $rowIndex
     * @return Integer[]
     */
    function getRow($rowIndex) {
        if($rowIndex == 0) return [1];

        $dp = [];
        $tmpdp = [1];
        for($i=1;$i<$rowIndex+1;$i++){
            $dp = [];
            for($j=0;$j<$i+1;$j++){
                if(isset($tmpdp[$j-1]) && isset($tmpdp[$j])){
                    $dp[$j] = $tmpdp[$j-1] + $tmpdp[$j];
                }else{
                    $dp[$j] = 1;
                }
            }

            $tmpdp = $dp;
        }

        return $dp;
    }
}


$a = new Solution();
$numRows = 3;
var_dump($a->getRow($numRows));
