<?php

class Solution {

    /**
     * @param Integer[] $digits
     * @return Integer[]
     */
    function plusOne($digits) {
        if(empty($digits)) return $digits;

        $j = count($digits)-1;
        for(;$j>=0;$j--){
            if($digits[$j] != 9){
                break;
            }
        }

        $res = [];
        if($j < count($digits) - 1){
            $i=0;
            for(;$i<$j;$i++){
                $res[] = $digits[$i];
            }

            if(isset($digits[$j])){
                $res[] = $digits[$i] + 1;
                $end = count($digits) - $i - 1;
                for($n = 0; $n<$end;$n++){
                    $res[] = 0;
                }
            }else{
                $res[] = 1;
                for($n = 0; $n < count($digits); $n++){
                    $res[] = 0;
                }
            }
        }else{
            $res = $digits;
            $res[count($res)-1] += 1;
        }

        return $res;
    }
}

$a = new Solution();
$digits = [1,2,0,9];
var_dump($a->plusOne($digits));

/*
题解：
从数组尾部获取最后一个为9的下标，然后按照如下规则处理
1、如果整个数组为 全部为9的，那么 新增一位为1，其余全部为0
2、如果中间的为9而不是最后一位为9，那么直接最后一位加1
3、如果从最后位连续好多位置均为9，那么获取最后一个为9的下标，将该下标的前一位加1，其余的均为0直接赋值即可。

*/
