<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function alternateDigitSum($n) {
        if($n == 0) return 0;

        $res = 0;
        $index = 0;
        while($n){
            $tmp = $n % 10;
            $n = intval(floor($n/10));

            if(($index % 2) == 0){
                $res += $tmp;
            }else{
                $res += (-1*$tmp);
            }
            $index++;
        }

        return -1*$res;
    }
}

$aa =  new Solution();
var_dump($aa->alternateDigitSum(521));

//犯了一个错误，判断是否是偶数还是奇数位置时使用了 /，而不是是 %, 除号只能是2的倍数，除数的结果不一定为0，只有余数才为0
