<?php

/*

给你长度相等的两个字符串 s1 和 s2 。一次 字符串交换 操作的步骤如下：选出某个字符串中的两个下标（不必不同），并交换这两个下标所对应的字符。

如果对 其中一个字符串 执行 最多一次字符串交换 就可以使两个字符串相等，返回 true ；否则，返回 false 。

 

示例 1：

输入：s1 = "bank", s2 = "kanb"
输出：true
解释：例如，交换 s2 中的第一个和最后一个字符可以得到 "bank"
示例 2：

输入：s1 = "attack", s2 = "defend"
输出：false
解释：一次字符串交换无法使两个字符串相等
示例 3：

输入：s1 = "kelb", s2 = "kelb"
输出：true
解释：两个字符串已经相等，所以不需要进行字符串交换
示例 4：

输入：s1 = "abcd", s2 = "dcba"
输出：false
 

提示：

1 <= s1.length, s2.length <= 100
s1.length == s2.length
s1 和 s2 仅由小写英文字母组成

*/

class Solution {
    /**
    * @param String $s1
    * @param String $s2
    * @return Boolean
    */
    function areAlmostEqual($s1, $s2) {
        if(strcmp($s1, $s2) == 0) return true;

        if(empty($s1) || empty($s2)) return false;

        if(strlen($s1) != strlen($s2)) return false;

        $arr1 = str_split($s1);
        $arr2 = str_split($s2);

        $res = true;
        $diff = [];
        $diffCnt = 0;
        for($i=0;$i<count($arr1);$i++){
            if($arr1[$i] != $arr2[$i]){
                if(empty($diff)){
                    $diff[] = $arr1[$i];
                    $diff[] = $arr2[$i];
                }else{
                    if($arr2[$i] != $diff[0] || $arr1[$i] != $diff[1]){
                        $res = false;
                        break;
                    }else{
                        $diff = [];
                        if($diffCnt >= 1) {
                            $res = false;
                            break;
                        }
                        $diffCnt++;
                    }
                }
            }
        }

        if(!empty($diff)) return false;

        return $res;
    }
}
