<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function fib($n) {
        if($n == 0) return 0;
        if($n == 1 || $n == 2) return 1;
        $result = 1;
        $next = 1;
        $pre = 0;
        for($i = 2; $i <= $n; $i++){
            $result = $pre + $next;
            $pre = $next;
            $next = $result;
        }

        return $result;
    }
}

/*
题目地址：
https://leetcode.cn/problems/fibonacci-number/
题解：
首先考虑递归，因为实现简单，而且只需要几行代码就可以实现，但是提交之后一直报错，可能LeetCode不想让我搞递归，遂使用非递归的模式，代码如上，思路如下：
下标从2开始，因为0和1 的返回值都是确定的，然后循环里面只要思路为： F(n) = F(n-1) + F(n-2)， pre为n-2, next 为n-1; 
首先计算出本次的 F(n)的值，然后 做赋值操作，原因是因为下一次的F(n)的计算 时 本次计算的 F(n) 变成了 F(n-1)， 如此类推，直接ac， 后续发现少了一个对于0 的边界值的判断，加上直接过
*/
