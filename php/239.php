<?php

/*
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
 *
 * */

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function maxSlidingWindow($nums, $k) {
        if(empty($nums) || count($nums) == 1) return $nums;

        $res = [];
        $j = 0;
        $tmp_max = PHP_INT_MIN;
        $maxindex = 0;
        for(;$j < $k; $j++){
            if($nums[$j] > $tmp_max){
                $tmp_max = $nums[$j];
                $maxindex = $j;
            }
        }

        $i = 1;
        $res[] = $tmp_max;
        for(;$j < count($nums);$j++,$i++){
            if($maxindex >= $i && $nums[$j] <= $tmp_max){            //在当前集合里面
                $res[] = $tmp_max;
            }else{
                list($len_index, $len_max) = $this->getMax($nums, $i, $i+$k-1);
                $tmp_max = $len_max;
                $maxindex = $len_index;
                $res[] = $tmp_max;
            }
        }

        return $res;
    }

    function getMax($nums, $i, $j){
        $max = PHP_INT_MIN;
        $index = -1;
        while($i<=$j){
            if(isset($nums[$i]) && $nums[$i] > $max){
                $max = $nums[$i];
                $index =$i;
            }
            $i++;
        }

        return [$index,$max];
    }
}

$nums = [1,-1];
$k = 1;
$a = new Solution();
var_dump($a->maxSlidingWindow($nums, $k));


/*
题解：这是一道hard的题目，我的解题思路是：
先查找前k个最大的，然后记录，滑动窗口每次往前移动，判断最大值是否在里面，不在则计算出最大值以及下标，在里面，则判断新加入的是否比最大值大，小则直接赋值

没有ac过去，超时，这种解法是最简单的解法，不是最优解

*/
