<?php

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
    function oddEvenList($head) {
        if(empty($head)) return null;

        $odd = $head;
        $evenHead = $odd->next;
        $even = $evenHead;
        while($even != null && $even->next != null){
            $odd->next = $even->next;
            $odd = $odd->next;
            $even->next = $odd->next;
            $even = $even->next;
        }

        $odd->next = $evenHead;
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

$aaa = [1,2,3,4,5,6,7,8];
$list = init($aaa);

$a = new Solution();
var_dump($a->oddEvenList($list));
