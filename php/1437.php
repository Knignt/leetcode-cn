<?php

/*
 *
 *
 * 给你一个由若干 0 和 1 组成的数组 nums 以及整数 k。如果所有 1 都至少相隔 k 个元素，则返回 True ；否则，返回 False 。



示例 1：



输入：nums = [1,0,0,0,1,0,0,1], k = 2
输出：true
解释：每个 1 都至少相隔 2 个元素。
示例 2：



输入：nums = [1,0,0,1,0,1], k = 2
输出：false
解释：第二个 1 和第三个 1 之间只隔了 1 个元素。
示例 3：

输入：nums = [1,1,1,1,1], k = 0
输出：true
示例 4：

输入：nums = [0,1,0,1], k = 1
输出：true


提示：

1 <= nums.length <= 10^5
0 <= k <= nums.length
nums[i] 的值为 0 或 1

 * */

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    function kLengthApart($nums, $k) {
        if(empty($nums) || $k <= 0) return true;

        $index = -999;          //处理当第一个是0的情况，
        $isOk = true;
        for($i=0;$i<count($nums);$i++){
            if($nums[$i] == 1){
                if($i - $index -1 < $k){        //需要减去本身
                    $isOk = false;
                    break;
                }else{
                    $index = $i;
                }
            }
        }

        return $isOk;
    }
}
