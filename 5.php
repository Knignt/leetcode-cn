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

        $dp = [];
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

        return substr($s, $resi, $resj - $resi + 1);
    }
}

$a = new Solution();
$s = "borcdubqiupahpwjizjjbkncelfazeqynfbrnzuvbhjnyvrlkdyfyjjxnprfocrmisugizsgvhszooktdwhehojbrdbtgkiwkhpfjfcstwcajiuediebdhukqnroxbgtvottummbypokdljjvnthcbesoyigscekrqswdpalnjnhcnqrarxuufzzmkwizptyvjhpfidgmskuaggtpxqisdlyloznkarxofzeeyvteynluofuhbllyiszszbwnsglqjkignusarxsjvctpgiwnhdufmfpanfwxjwlmhgllzsmdpncbwnsbdtsvrjybynifywkulqnzprcxockbhrhvnwxrkvwumyieazclcviumvormynbryaziijpdinwatwqppamfiqfiojgwkvfzyxadyfjrgmtttvlgkqghgbcfhkigfojjkhskzenlpasyozcsuccdxhulcubsgomvcrbqwakrraesfifecmoztayrdjicypakrrneulfbqqxtrdelggedvloebqaztmfyfkhuonrognejxpesmwgnmnnnamvkxqvzrgzdqtvuhccryeowywmbixktnkqnwzlzfvloyqcageugmjqihyjxhssmhkfzxpzxmgpjgsduogfolnkkqizitbvvgrkczmojxnabruwwzxxqcevdwvtiigwckpxnnxxxdhxpgomncttjutrsvyrqcfwxhexhaguddkiesmfrqfjfdaqbkeqinwicphktffuwcazlmerdhraufbpcznbuipmymipxbsdhuesmcwnkdvilqbnkaglloswcpqzdggnhjdkkumfuzihilrpcvemwllicoqiugobhrwdxtoefynqmccamhdtxujfudbglmiwqklriolqfkknbmindxtljulnxhipsieyohnczddabrxzjmompbtnnxvivmoyfzfrxlufowtqzojfclmatthjtbhotdoheusnpirwicbtyrcuojsdmfcx";
var_dump($a->longestPalindrome($s));


