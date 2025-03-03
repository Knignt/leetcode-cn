<?php

/*
 *
 * 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


进阶：

尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 * */

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return NULL
     */
    function rotate(&$nums, $k) {
        if($k < 0) return $nums;

        $k = $k % count($nums);             //当传入k的值大于数组长度的话，需要取余
        $k = count($nums) - $k;             //这里是向右轮转，所以就是后$k位向前移动，对于前面而言，就是 count($nums) - $k 个移动位置。


        $arr1 = array_reverse(array_splice($nums, 0, $k));
        $arr2 = array_reverse($nums);
        $nums = array_reverse(array_merge($arr1, $arr2));
    }
}

$a = new Solution();
$nums = [1,2,3,4,5,6,7];
$k = 3;
$a->rotate($nums, $k);
var_dump($nums);
