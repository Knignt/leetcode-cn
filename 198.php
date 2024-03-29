<?php

/**

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。



示例 1：

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
偷窃到的最高金额 = 1 + 3 = 4 。
示例 2：

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
偷窃到的最高金额 = 2 + 9 + 1 = 12 。


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 400

 **/

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        if(empty($nums)) return 0;

        //dp[i] = max(dp[i-1], dp[i-2] + nums[i])
        $dp = [];
        $dp[0] = $nums[0];
        $dp[1] = $nums[1] > $nums[0] ? $nums[1] : $nums[0];
        for($i =2; $i < count($nums);$i++){
            $dp[$i] = max($dp[$i-1], $dp[$i-2] + $nums[$i]);
        }

        return $dp[count($nums)-1];
    }
}

$a = new Solution();
$nums = [2,1,1,2];
var_dump($a->rob($nums));




/**
题解：这里最主要的点在于
1、dp[i-1]表示前面抢劫的和的最大值
2、dp[0]和dp[1] 分别表示当数组长度为0或者1 的时候，抢劫的最大值，所以 dp[1]的赋值应该是 前两个的最大值
3、dp[i] = max(dp[i-1], dp[i-2]+nums[i])  这里理解如下：
    a、如果不抢劫这个屋子，那么 dp[i-1]
    b、如果抢劫这个屋子，按照规则，需要dp[i-2]和当前值
    这两个比较大小，因为总是需要的是最大值
*/
