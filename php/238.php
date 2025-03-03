<?php

/**
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：

2 <= nums.length <= 105
-30 <= nums[i] <= 30
保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内

 **/

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function productExceptSelf($nums) {
        if(empty($nums)) return [];

        $left = $right = [];
        $left[0] = 1;
        for($i= 1;$i<count($nums);$i++){
            $left[$i] = $left[$i-1]*$nums[$i-1];
        }

        $right[count($nums)-1] = 1;
        for($i = count($nums)-2;$i>=0;$i--){
            $right[$i] = $right[$i+1]*$nums[$i+1];
        }

        $res = [];
        for($i=0;$i<count($nums);$i++){
            $res[$i] = $left[$i] * $right[$i];
        }

        return $res;
    }
}

$a = new Solution();
$nums = [1,2,3,4];
var_dump($a->productExceptSelf($nums));

/***
 * 解题思路：
 * 根据前缀和思想，首先我们想，要求一个除过该数字以外的其余数字的积该如何求，左边的积乘以右边积
 * 左边的积该如何求，从第一个开始，只计算除过当前i的前面的积，那么就是从0开始，赋值为1，这样下标为0的就不会被计算进去，然后i-1乘以前面的积
 * 右边的积计算的方法类似，这样的话，只需要O(n)的方式来计算积即可。
 **/
