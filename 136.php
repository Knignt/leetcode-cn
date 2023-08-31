<?php


/*
 *
 * 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。



示例 1 ：

输入：nums = [2,2,1]
输出：1
示例 2 ：

输入：nums = [4,1,2,1,2]
输出：4
示例 3 ：

输入：nums = [1]
输出：1


提示：

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
除了某个元素只出现一次以外，其余每个元素均出现两次。
 *
 * */


class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function singleNumber($nums) {
        if(count($nums) == 1) return $nums[0];
        $res = $nums[0];
        for($i=1;$i<count($nums);$i++){
            $res ^= $nums[$i];
        }

        return $res;
    }
}
$nums = [-1];
$a = new Solution();
var_dump($a->singleNumber($nums));


/*
 * 主要要知道亦或运算满足 交换律 和 结合律，即
 * a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b
 *
 * a ^ b ^ a = a ^ a ^ b = 0 ^ b = b
 * */
