<?php

/*
 *给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。



示例 1：


输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]


提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
 *
 *
 *
 * */

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
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
     * @return ListNode
     */
    function swapPairs($head) {
        if(empty($head) || $head->next == null) return $head;

        $a = $head;
        $b = $head->next;
        $cur = new ListNode(0);
        $res = $cur;
        $cur->next = $head;
        while($b != null){
            $cur->next = $b;
            $a->next = $b->next;
            $b->next = $a;
            $cur = $a;
            $b = $a->next->next;
            $a= $a->next;
        }

        return $res->next;
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

$aaa = [1,2,3,4,5,6,7,8];
$list = init($aaa);

$bbb = new Solution();
$res = $bbb->swapPairs($list);
var_dump($res);

/*
题解：
这里主要的做法是需要新增一个虚拟的头结点，然后有一个cur来处理交换的那个节点的指向，比如头指针和第二个节点，等交换完毕，应该把头指针指向第二个节点，因为第一个节点已经被交换到第二个节点的后面了。
*/
