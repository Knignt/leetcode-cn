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
        /*if(empty($list2) && empty($list1)) return false;
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

        return $head;*/
        
        if(empty($list1)) return $list2;
        if(empty($list2)) return $list1;

        if($list1->val <= $list2->val){
            $list1->next = $this->mergeTwoLists($list1->next, $list2);
            return $list1;
        }else{
            $list2->next = $this->mergeTwoLists($list1, $list2->next);
            return $list2;
        }
    }
}
//leetcode submit region end(Prohibit modification and deletion)

/*
合并链表，大体思路：
1、基本方法，也是我目前提交的方案，首先处理表头，链处理，以及最后剩余部分的处理，简单思路，不做具体的说明在此
2、这道题目是 递归 类别里面的题目，递归代码我还没看懂，看懂再做更新
首先贴一下大佬的解释：
https://leetcode.cn/problems/merge-two-sorted-lists/solutions/103891/yi-kan-jiu-hui-yi-xie-jiu-fei-xiang-jie-di-gui-by-/

如下为个人理解：
1、递归首先为 自己调用自己，这是基本用法
2、主要点为两个：第一，搞清楚边界值，递归不能无限循环下去；第二，搞清楚这一层需要做什么事情

对于这道题而言，我一直不理解的点在于，如何合并
目前的理解是，递归的方案没有申请多余的空间，只是将小的值放在了前面，即小的值后面拼接的是大的值的(这部分可能是一整个的链表，而不仅仅是单个值)，
与此同时将另外小的那个next当作下一层的另外一条链的链表头做输入，以此来进行下一层继续运算。
*/
