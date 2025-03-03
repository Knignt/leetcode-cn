<?php
/*
 * 给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105


进阶：

如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 * */


class Solution {

    /**
     * @param Integer $target
     * @param Integer[] $nums
     * @return Integer
     */
    function minSubArrayLen($target, $nums) {
        if(empty($nums)) return 0;

        $tmp_sum = 0;
        $len = PHP_INT_MAX;
        $start = $end = 0;
        for(;$end < count($nums);$end++){
            if($tmp_sum < $target){
                $tmp_sum += $nums[$end];
            }

            while($tmp_sum >= $target){
                if($len > ($end - $start + 1)){
                    $len = ($end - $start + 1);
                }
                $tmp_sum -= $nums[$start];
                $start++;
            }
        }

             if($len == PHP_INT_MAX) $len = 0;

        return $len;
    }
}

$a = new Solution();
$target = 11;
$nums = [1,2,3,4,5];
var_dump($a->minSubArrayLen($target, $nums));


/*
题解：
滑动窗口的思路我有，但是主要问题在于里面的while循环，题目要求找最小，那么其实只要满足 tmp_sum >= target，就需要一直把滑动窗口的开始点往做移动，找到最小的长度

开始给len赋值的时候，我给的值是 0，这样就会导致最短的永远是0，需要赋值为 PHP_INT_MAX，这样len有变化就会赋值

还有一个case没过，就是如果最后没找到最短的长度，需要价格判断，返回0

*/
