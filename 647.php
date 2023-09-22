<?php

/**
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。



示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"


提示：

1 <= s.length <= 1000
s 由小写英文字母组成

 **/

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function countSubstrings($s) {

        if(empty($s)) return 0;
        if(strlen($s) == 1) return 1;

        $len = strlen($s);
        $cnt  = 0;
        $dp = [];
        for($i=0;$i<$len;$i++){
            for($j=0;$j<($len);$j++){
                if($j == $i){
                    $dp[$i][$j] = true;
                }else{
                    $dp[$i][$j] = false;
                }
            }
        }

        for($i=$len-2;$i>=0;$i--){
            for($j=$i+1;$j<($len);$j++){
                if($dp[$i + 1][$j - 1] && $s[$i] == $s[$j]){
                    $dp[$i][$j] = true;
                }

                if($i+1 == $j && $s[$i] == $s[$j]){
                    $dp[$i][$j] = true;
                }
            }
        }

        for($i=0;$i<$len;$i++){
            for($j=0;$j<$len;$j++){
                if($dp[$i][$j]) $cnt++;
            }
        }

        return $cnt;
    }
}

$a = new Solution();
$s = "abc";
var_dump($a->countSubstrings($s));

/**
题解：
首先需要想到的一点是：
如果某一个子串s,左右两边分别新增字符a和b，如果ab相等，那么新的字符串也是回文的

那么dp数组怎么赋值呢？
dp数组应该是一个二维数组，它表示从字符i开始到字符j结束，这段子串是否是回文串

 * 初始化， i == j 的是 true，其余为false

 * 循环阶段，首先我i从0开始的，后来发现 dp[i+1][j-1]这个还没有计算啊，必须先计算这个，所以，我就从len-2开始了，因为len-1不需要计算，最后一行之后一个true
 * 其次，我漏掉了一个情况，就是在 子串长度为2的情况下 ,即 i + 1 == j，并且 s[i] == s[j],那么也为true
 * 为什么这种情况为true，如果按照递推公式来计算，那么 dp[i+1][j-1] 会出现 i > j 的情况，从dp数组里面获取为false，所以，需要计算一遍
 **/


