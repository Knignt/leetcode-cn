<?php

/**
给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。


示例 1：

输入：words = ["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：words = ["cool","lock","cook"]
输出：["c","o"]


提示：

1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] 由小写英文字母组成
 **/

class Solution {

    /**
     * @param String[] $words
     * @return String[]
     */
    function commonChars($words) {
        if(empty($words)) return [];


        //这里先把每一个元素的个数计算出来
        $hash = $res = [];
        foreach($words as $row){
            $row1 = str_split($row);
            $hash[$row] = [];
            foreach($row1 as $item){
                if(!isset($hash[$row][$item])){
                    $hash[$row][$item] = 1;
                }else{
                    $hash[$row][$item]++;
                }
            }
        }


        //找出每一个元素的相同的，并且找出每一个单词的最小值，这里最小值的意义在于，比如在 cooll 和 coolll 以及cool 这三个里面，l这个单词应该公共的只有一次， 共有的是 c o o l 这四个单词
        foreach($hash[$words[0]] as $key => $value){
            $min = PHP_INT_MAX;
            for($i=1;$i<count($words);$i++){
                if(isset($hash[$words[$i]][$key])){
                    $min = min($min, $value, $hash[$words[$i]][$key]);
                }else{
                    $min = 0;
                    break;
                }
            }

            if($min > 0){
                while($min>0){
                    $res[] = $key;
                    $min--;
                }
            }


        }

        return $res;
    }
}

$a = new Solution();
$words = ["acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"];
//$words = ["cool","lock","cook"];
var_dump($a->commonChars($words));


