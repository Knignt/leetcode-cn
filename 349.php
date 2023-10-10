<?php
/**
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
 **/

class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    function intersection($nums1, $nums2) {
        if(empty($nums1) || empty($nums2)) return [];

        $hash = [];

        $common = [];
        for($i=0;$i<count($nums1);$i++){
            $hash[$nums1[$i]] = 1;
        }

        for($j=0;$j<count($nums2);$j++){
            if(isset($hash[$nums2[$j]]) && !in_array($nums2[$j], $common)){
                $common[] = $nums2[$j];
            }
        }

        return $common;
    }
}

$a = new Solution();
$nums1 = [1,2,2,1];
$nums2 = [2,2];
var_dump($a->intersection($nums1, $nums2));
