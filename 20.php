<?php

class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        if(empty($s)) return false;

        $temp = [
            ')' => '(',
            '}' => '{',
            ']' => '['
        ];

        $s = str_split($s);
        if(count($s) %2 != 0) return false;

        $res = true;
        $stack = [];
        foreach($s as $item){
            if(isset($temp[$item])){
                $pop = array_pop($stack);
                if($pop != $temp[$item]){
                    $res = false;
                    break;
                }
            }else{
                $stack[] = $item;
            }
        }

        if(!empty($stack)) return false;

        return $res;
    }
}

$a = new Solution();
$s = "{[]}";
var_dump($a->isValid($s));
