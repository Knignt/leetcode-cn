<?php

/**
给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。



示例 1：

输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。
示例 2：

输入：nums = [2,2,3,3,3,4]
输出：9
解释：
删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。


提示：

1 <= nums.length <= 2 * 104
1 <= nums[i] <= 104

 **/

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function deleteAndEarn($nums) {
        if(empty($nums)) return 0;

        $max = PHP_INT_MIN;
        for($i=0;$i<count($nums);$i++){
            $max = max($max, $nums[$i]);
        }

        $hash = [];
        for($i=0;$i<count($nums);$i++){
            if(!isset($hash[$nums[$i]])) $hash[$nums[$i]]=0;
            $hash[$nums[$i]]++;
        }
        for($i=0;$i<=$max;$i++){
            if(!isset($hash[$i])) $hash[$i]=0;
        }
        ksort($hash);

        $dp = [];
        $dp[0] = $hash[0];
        $dp[1] = max($dp[0], $hash[1]);
        for($i=2;$i<count($hash);$i++){
            $dp[$i] = max($dp[$i-1], $dp[$i-2] + $hash[$i]*$i);
        }

        return $dp[count($hash)-1];
    }
}

/**
这道题和打家截舍最大的差别在于：
 * 需要进行一次转换，因为nums[i]-1和nums[i]+1 需要删除，其实也就意味着，在自然增长的数列中，左右邻居不能抢，剩下的思路和打家劫舍一致
 * 另外，当前的分数应该个数乘以分值.因为是删除所有该值。
 **/


$a = new Solution();
$nums = [3,4,2];
var_dump($a->deleteAndEarn($nums));
