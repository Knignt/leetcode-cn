<?php

class Solution {

    /**
     * @param String[] $s
     * @return NULL
     */
    function reverseString(&$s) {
        if(empty($s)) return [];

        for($i=0,$cnt = count($s)-1;$cnt>$i;$cnt--,$i++){
            $tmp=$s[$i];
            $s[$i] = $s[$cnt];
            $s[$cnt] = $tmp;
        }

        return $s;
    }
}

$a = new Solution();
$s = ["h","e","l","l","o"];
var_dump($a->reverseString($s));
