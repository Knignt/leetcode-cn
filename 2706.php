<?php

/***
给你一个整数数组 prices ，它表示一个商店里若干巧克力的价格。同时给你一个整数 money ，表示你一开始拥有的钱数。

你必须购买 恰好 两块巧克力，而且剩余的钱数必须是 非负数 。同时你想最小化购买两块巧克力的总花费。

请你返回在购买两块巧克力后，最多能剩下多少钱。如果购买任意两块巧克力都超过了你拥有的钱，请你返回 money 。注意剩余钱数必须是非负数。



示例 1：

输入：prices = [1,2,2], money = 3
输出：0
解释：分别购买价格为 1 和 2 的巧克力。你剩下 3 - 3 = 0 块钱。所以我们返回 0 。
示例 2：

输入：prices = [3,2,3], money = 3
输出：3
解释：购买任意 2 块巧克力都会超过你拥有的钱数，所以我们返回 3 。


提示：

2 <= prices.length <= 50
1 <= prices[i] <= 100
1 <= money <= 100
 ***/


class Solution {

    /**
     * @param Integer[] $prices
     * @param Integer $money
     * @return Integer
     */
    function buyChoco($prices, $money) {
        if(empty($prices) || $money <= 0) return $money;

        sort($prices);
        $res = $money;
        $cnt = 0;
        for($i=0;$i<count($prices);$i++){

            //判断当前是否2块巧克力是否买够了
            if($cnt >= 2){
                break;
            }

            //如果不够去买，那么就退出
            if($res - $prices[$i] < 0){
                break;
            }

            $res -= $prices[$i];
            $cnt++;
        }

        if($cnt >= 2){
            return $res;
        }else{
            return $money;
        }
    }
}

$a = new Solution();
$prices = [3,2,3];
$money = 3;
var_dump($a->buyChoco($prices, $money));
