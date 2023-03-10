<?php
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 3013 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

//Definition for a singly-linked list.
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

function printList($head){
        //$i=0;
    while($head){
                var_dump($head->val);
                $head = $head->next;
                //$i++;
                //if($i> 10) break;
    }
}

class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        if(empty($head)) return null;

        $a = $b = $c = $head;
        while($c){
            $c = $b->next;
            $b->next = $a;
            $a = $b;
            $b = $c;
        }
                $head->next = null;

        return $a;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

$a = new Solution();
$b = [1,2,3,4,5,6,7,8,9];

$list = $temp = new ListNode($b[0]);
for($i=1;$i<count($b);$i++){
    $tmp = new ListNode($b[$i]);
    $temp->next = $tmp;
        $temp = $tmp;
}

printList($list);
$d = $a->reverseList($list);
printList($d);

/*
思路分析：
简单题目
三个指针，一个永远指向最后返回的指针头，一个指向被迁移的位置，另外一个指向被迁移位置的next

写的时候，最后自己忘记把head->next置空导致回环
*/
