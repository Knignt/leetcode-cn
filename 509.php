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