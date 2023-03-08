<?php
//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//
//
//提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 1711 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement($nums) {
        if(empty($nums)) return 0;

        $candidates = $nums[0];
        $count = 0;
        foreach ($nums as $num){
            if($count == 0){
                $candidates = $num;
            }

            if($candidates == $num){
                $count++;
            }else{
                $count--;
            }
        }

        return $candidates;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

$a = new Solution();
$list = [3,2,3];
var_dump($a->majorityElement($list));

/*
对于这个题目，个人思路是使用分治的方法，but,我没写出来，然后对照别人的题解，写了一个 摩尔投票的算法来处理，
关于摩尔算法的相关资料如下：
https://zhuanlan.zhihu.com/p/346342647

分治等我学会了再来装b，不好意思~
*/
