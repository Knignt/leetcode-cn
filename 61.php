<?php

/*给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

 

示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：


输入：head = [0,1,2], k = 4
输出：[2,0,1]
 

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109
*/



class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution {

    /**
     * @param ListNode $head
     * @param Integer $k
     * @return ListNode
     */
    function rotateRight($head, $k) {
        if(empty($head) || $k == 0) return $head;

        $arr = [];
        $res = $head;
        $index = 0;
        while($head!= null){
            ///$arr[$index++] = $head->val;
            $head = $head->next;
            $index++;
        }
        if($index == 1) return $res;

        $head = $res;

        $k = $k % $index;
        $k = $index - $k;
        if($index == $k) return $head;

        $cnt = 0;
        while($cnt < $k-1){
            $head = $head->next;
            $cnt++;
        }
        $tmpHead = $head;
        $head = $head->next;
        $tmpHead->next = null;

        $tmp = $head;
        while($tmp->next != null){
            $tmp = $tmp->next;
        }
        $tmp->next = $res;

        return $head;
    }
}

function init($arr){

    $head = null;
    $list = null;
    foreach($arr as $row){
        $tmp = new ListNode($row);
        if(empty($head)){
            $head = $tmp;
            $list = $tmp;
        }else{
            $head->next = $tmp;
            $head = $head->next;
        }
    }

    return $list;
}

$aaa = [1,2];
$val = 2;
$list = init($aaa);

$a = new Solution();
var_dump($a->rotateRight($list, $val));

/*题解：
1、把需要旋转的两部分拆除出来，然后合并在一起即可

需要注意的细节点：
1、链表传入的旋转位数大于链表长度
2、旋转的位数和长度一致，则没必要旋转，因为旋转完毕的结果也是原链表
*/
