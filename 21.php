<?php
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 2965 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
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
     * @param ListNode $list1
     * @param ListNode $list2
     * @return ListNode
     */
    function mergeTwoLists($list1, $list2) {
        if(empty($list2) && empty($list1)) return false;
        if(empty($list1)) return $list2;
        if(empty($list2)) return $list1;

        //表头处理
        $res = $head = new ListNode();
        if($list1->val >= $list2->val){
            $res->val = $list2->val;
            $list2 = $list2->next;
        }else{
            $res->val = $list1->val;
            $list1 = $list1->next;
        }

        //链处理
        while (true){
            if(empty($list2) || empty($list1)) break;

            if($list2->val <= $list1->val){
                $res->next = $list2;
                $list2 = $list2->next;
            }else{
                $res->next = $list1;
                $list1 = $list1->next;
            }
            $res = $res->next;
        }

        //最后多余的处理
        if(empty($list2)){
            $res->next = $list1;
        }

        if(empty($list1)){
            $res->next = $list2;
        }

        return $head;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

/*
合并链表，大体思路：
1、基本方法，也是我目前提交的方案，首先处理表头，链处理，以及最后剩余部分的处理，简单思路，不做具体的说明在此
2、这道题目是 递归 类别里面的题目，递归代码我还没看懂，看懂再做更新
*/
