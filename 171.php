<?php

/**
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


示例 1:

输入: columnTitle = "A"
输出: 1
示例 2:

输入: columnTitle = "AB"
输出: 28
示例 3:

输入: columnTitle = "ZY"
输出: 701


提示：

1 <= columnTitle.length <= 7
columnTitle 仅由大写英文组成
columnTitle 在范围 ["A", "FXSHRXW"] 内
 **/

class Solution {

    /**
     * @param String $columnTitle
     * @return Integer
     */
    function titleToNumber($columnTitle) {
        if (empty($columnTitle)) return 0;

        $res = 0;
        $len = strlen($columnTitle);
        for($i=0;$i<$len;$i++){
            $res =  $res*26 + (ord($columnTitle[$i]) - ord('A')+1);
        }
        return ($res > 2147483647) ? 2147483647:$res;
    }
}

$a = new Solution();
$columnTitle = "AAA";
var_dump($a->titleToNumber($columnTitle));
