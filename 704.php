<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function search($nums, $target) {
        if(empty($nums)) return -1;

        $index = -1;
        $i = 0;
        $j = count($nums)-1;
        while($i <= $j){
            $mid = floor(($i + $j) / 2);

            if($nums[$mid] > $target){
                $j = $mid-1;
            }elseif($nums[$mid] < $target){
                $i = $mid+1;
            }else{
                $index = $mid;
                break;
            }

        }
        return $index;
    }
}

$a = new Solution();
$nums = [2,5];
$target = 5;
var_dump($a->search($nums, $target));
