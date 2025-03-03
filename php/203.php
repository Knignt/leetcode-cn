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
     * @param Integer $val
     * @return ListNode
     */
    function removeElements($head, $val) {
        if(empty($head)) return null;

        $pre = $res = $head;
        $head = $head->next;
        while(!empty($head)){
            if($head->val == $val){
                $pre->next = $head->next;
                //continue;
            }else{
                $pre = $head;
            }
            $head = $head->next;
        }

        if($res->val == $val){
            $res = $res->next;
        }

        return $res;
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

$aaa = [1,2,6,3,4,5,6];
$val = 6;
$list = init($aaa);

$bbb = new Solution();
$res = $bbb->removeElements($list, $val);
var_dump($res);
