<?php
/*
 *
 *给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

 

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
 

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
* */

class Solution {

    /**
     * @param String[] $strs
     * @return String[][]
     */
	function groupAnagrams($strs) {
		$res = [];
		foreach($strs as $row){
			$tmp =  $row;
			$row = str_split($row);
			sort($row, SORT_STRING);
			$row = implode($row);
			$res[$row][] = $tmp;
		}

		return array_values($res);
	}
}


$aa = new Solution();
$b = ["eat", "tea", "tan", "ate", "nat", "bat"];
var_dump($aa->groupAnagrams($b));


/*
解题思路：
开始的解题思路是，将每一个item的全排列计算出来，再去搜寻，在数组的内的放到同一个里面，但是这样需要求全排列，
所以，考虑转变思路，主要思路为，不管怎么样，对于一个字符串，不管怎么变更字符，字符串内的字符仍旧不变，不管是 abc 还是bca，均是由a、b、c 这三个字母组成的字符串，
故再次提出如下两种思路，并给出其中一种思路的解答方式：
1、将每个字符串内的字母按照字典序排好，这样的话，字母异位词的值是相同，可以归纳到同一个数组里面
2、寻找一个f(x)，另 f("abc") = f("bca")，即可求解
*/
