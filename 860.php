<?php


class Solution {

    /**
     * @param Integer[] $bills
     * @return Boolean
     */
    function lemonadeChange($bills) {
        if(empty($bills)) return false;

        $arr = [];
        $res = true;
        foreach($bills as $cost){
            if($cost == 5){
                $arr[] = $cost;
                continue;
            }

            $tmp = 0;
            $isCan = false;
            $index = 0;
            $count = count($arr)-1;
            for(;$count>0;$count--){
                $tmp += $arr[$count];
                $index++;
                if($tmp == $cost - 5){
                    $isCan = true;
                    $arr = array_slice($arr, 0, $count);
                    $arr[] = $cost;
                    var_dump($arr, "++++++++++++++");
                    break;
                }elseif($tmp > $cost-5){
                    echo 77777777777;
                    $tmp -= $arr[$count];
                }
            }


            if(!$isCan){
                echo 2333;
                $res = false;
                break;
            }
        }

        return $res;
    }
}

$aa = new Solution();
$bills = [5,5,5,5,10,5,10,10,10,20];
var_dump($aa->lemonadeChange($bills));
