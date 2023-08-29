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



/*
题解：
这里的主要难点在于，如何循环，以及对于边界的把控（实在抱歉，我有点菜，所以，看了别人的题解）

循环：
每次循环分为四步
1、从左到右
2、从上至下
3、从有至左
4、从下至上


边界的把控，自己之前的想法是使用一个当前运行循环次数来进行限制，后来对照别人的代码才知道自己错误的点在于每次对于 i 和 j 的值进行重置，
所以这里，startx和starty可以使用times来代替

*/

