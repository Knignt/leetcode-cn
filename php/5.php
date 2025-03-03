<?php

/**
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
 **/

class Solution {

    /**
     * @param String $s
     * @return String
     */
    function longestPalindrome($s) {
        if(empty($s)) return 0;
        if(strlen($s) == 1) return $s;

        /*
         使用动态规划dp数组形式，时间复杂度O(n^2)，空间复杂度O(n^2)，对于测试实例里面的例子，会超时
         * $dp = [];
        $len = strlen($s);
        for($i=0;$i<$len;$i++){
            for($j=0;$j<$len;$j++){
                if($i==$j) $dp[$i][$j] = true;
            }
        }

        $maxLen = PHP_INT_MIN;
        $resi =  $resj = 0;
        for($i=$len-2;$i>=0;$i--){
            for($j=$i;$j<$len;$j++){
                if($j <= $i) continue;
                if($i==$j-1 && $s[$i] == $s[$j]) $dp[$i][$j] = true;
                if($dp[$i+1][$j-1] && $s[$i] == $s[$j]) $dp[$i][$j] = true;

                if($dp[$i][$j] && $j - $i > $maxLen){
                    $resi = $i;
                    $resj = $j;
                    $maxLen = $j - $i;
                }
            }
        }

        return substr($s, $resi, $resj - $resi + 1);*/


        /**
        下面的这种算法，使用中心扩展法，
         即单个单词以及两个单词为基准，左右扩展，如果新扩展的是相等的字符，则组成的为回文串
         对于如果两个单词不一样怎么办？在extends里面进去，对于i 和j 是否一致进行了判断，也就是说会先判断两个单词为基础的是否为回文串，这个不用担心
         另外，这里为什么只考虑两个单词和一个单词为基准，而不只考虑一个单词呢？
         举例而言，比如abba这样的单词，如果只根据b单词来扩展，那么其实第一次扩展就会失败，第一次扩展时 s[i] 为a, s[j] 为b，不是回文串
         *！！！这里需要注意，两个单词扩展时不存在 两个单词不一致的情况，即中心扩展词两个肯定一致，因为举不出来例子
         **/
        $len = strlen($s);
        $resi = $resj = -1;
        $maxLen = PHP_INT_MIN;
        for($i=0;$i<$len;$i++){
            $this->extend($s, $i, $i, $resi, $resj, $maxLen, $len);
            $this->extend($s, $i, $i+1, $resi, $resj, $maxLen, $len);
        }

        return substr($s, $resi, $resj - $resi + 1);
    }

    function extend($s, $i, $j, &$resi, &$resj, &$maxLen, $len){
        while($j<$len && $i >= 0 && $s[$i] == $s[$j]){
            if($j - $i +1 > $maxLen){
                $maxLen = $j - $i + 1;
                $resi = $i;
                $resj = $j;
            }
            $i--;
            $j++;
        }
    }
}

$a = new Solution();
$s = "aabb";
var_dump($a->longestPalindrome($s));


