<?php

/*

给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

 

示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]
 

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/

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
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function deleteDuplicates($head) {
        if(empty($head)) return [];

        $res = $head;
        while(!empty($head)){
            $tmp = $head;
            while(1){
                if(empty($head)){
                    $tmp->next=null;
                    break;
                }
                if($tmp->val == $head->val){
                    $head = $head->next;
                }else{
                    $tmp->next = $head;
                    break;
                }
            }
        }

        return $res;
    }
}
