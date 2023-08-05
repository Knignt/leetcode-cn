<?php

class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isAnagram($s, $t) {
        if(empty($s) || empty($t)) return false;
        $s = str_split($s);
        $t = str_split($t);
        if(count($s) != count($t)) return false;

        $tmps = $tmpt = [];
        foreach($s as $a){
            $tmps[$a]++;
        }

        foreach($t as $b){
            $tmpt[$b]++;
        }

        $ret = true;
        foreach($tmps as $k=>$v){
            if(!isset($tmpt[$k]) || $tmpt[$k] != $v){
                $ret = false;
                break;

            }
        } 

        return $ret;       
    }
}
