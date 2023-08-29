<?php

/*
 *给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。



示例 1：


输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 20
 *
 *
 * */

class Solution {

    /**
     * @param Integer $n
     * @return Integer[][]
     */
    function generateMatrix($n) {
        if($n<= 1) return [[1]];

        $res = [];
        $cnt = ($n/2);
        $i = $j = 0;
        $num = 1;
        $times = 0;

        $startx = $starty = 0;
        while($cnt > 0){
            $i = $startx;
            $j = $starty;

            for(;$j<$n-$times-1;$j++){
                $res[$i][$j] =  $num++;
            }

            for(;$i<$n-$times-1;$i++){
                $res[$i][$j] =  $num++;
            }

            for(;$j>$starty;$j--){
                $res[$i][$j] =  $num++;
            }

            for(;$i>$startx;$i--){
                $res[$i][$j] =  $num++;
            }

            $cnt--;
            $times++;

            $startx++;
            $starty++;
        }

        if($n %2 != 0){
            $res[$n/2][$n/2] = $num;
        }

        for($i=0;$i<$n;$i++){
            ksort($res[$i]);
        }
        return $res;
    }
}

$a = new Solution();
$n = 3;
var_dump($a->generateMatrix($n));
