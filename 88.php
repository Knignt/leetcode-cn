<?php

class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer $m
     * @param Integer[] $nums2
     * @param Integer $n
     * @return NULL
     */
    function merge(&$nums1, $m, $nums2, $n) {
        if($m == 0){
            $nums1 = $nums2;
            return;
        }

        if($n == 0){
            return;
        }

        $i = $m -1;
        $j = count($nums2)-1;
        $index = count($nums1)-1;
        while($i>=0 && $j >= 0){
            if($nums1[$i] > $nums2[$j]){
                $nums1[$index--] = $nums1[$i--];
            }else{
                $nums1[$index--] = $nums2[$j--];
            }
        }

        if($j >= 0){
            while($j >= 0){
                $nums1[$index--] = $nums2[$j--];
            }
        }

        if($i >= 0){
            while($i >= 0){
                $nums1[$index--] = $nums1[$i--];
            }
        }
    }
}

$a = new Solution();
$nums1 = [-1,0,0,3,3,3,0,0,0];
$m = 6;
$nums2 = [1,2,2];
$n = 3;
$a->merge($nums1, $m, $nums2, $n);
var_dump($nums1);

/*
题解思路：
开始没思路，看了题解，每一个给的值都是有效果的，比如给的m和n，主要思路为，在两个数组有序的前提下，从两个数组尾部往头部比大小，大的放在数组后面，并且该数组的下标往前挪，
我的问题在，刚开始在 while那边只想着，$i >= 0 就可以了啊，后面再加一个 对于 $j 的处理就OK， 但是忽略了一个比如，$j结束了，但是$i没结束，那么会陷入到死循环，所以修改为 $i && $j
然后后面新增 $j 和 $i 分别不为0的单独处理
*/
