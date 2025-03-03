<?php

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0


提示：

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
 **/

class Solution {

    /**
     * @param Integer[] $coins
     * @param Integer $amount
     * @return Integer
     */
    function coinChange($coins, $amount) {
        //边界值判定
        if(empty($coins) || $amount <= 0) return 0;



        $dp = [];       //装满背包的最小物品数 j 表示当前背包的容量
        $dp[0] = 0;
        for($i=1;$i<=$amount;$i++){
            $dp[$i] = PHP_INT_MAX;
        }

        for($i=0;$i<count($coins);$i++){                    //物品
            for($j=$coins[$i];$j<=$amount;$j++){             //背包       (包含空间为amount的个数)
                $dp[$j] = min($dp[$j-$coins[$i]]+1, $dp[$j]);
            }
        }

        if($dp[$amount] == PHP_INT_MAX) return -1;

        return $dp[$amount];
    }
}


$a = new Solution();
//$coins = [186,419,83,408];
$coins = [2];
//$amount = 6249;
$amount = 3;
var_dump($a->coinChange($coins, $amount));

/**
解题思路:
    划为子问题，背包大小一直为amount，每次给里面放 coins[i]体积的物品，dp[i]存储最小的次数
    双层循环理解 是 拿着 物品在背包里面去放，然后在dp里存储最小次数
    为什么需要 等于 amount，这里主要是因为 amount是背包大小，不可以比amount小
 **/
