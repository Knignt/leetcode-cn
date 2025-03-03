<?php

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。



示例 1：


输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1
 **/

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function numTrees($n) {
        if($n <= 2) return $n;

        //dp 数组为二维数组，i 为当前节点数 ,j 表示 以j为 根节点时共有多少种方案
        //递推公式: dp[i] += dp[j - 1] * dp[i - j]     //这里为什么需要减一，因为有一个root节点
        $dp = [];
        $dp[0] = 1;
        for($i=1;$i<=$n;$i++){
            $dp[$i] = 0;
            for($j=1;$j<=$i;$j++) {
                $dp[$i] += $dp[$j - 1] * $dp[$i - $j];
            }
        }
        return $dp[$n];
    }
}
$a = new Solution();
$n = 5;
var_dump($a->numTrees($n));

