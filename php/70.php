<?php
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
// 示例 2：
//
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//
//
// 提示：
//
//
// 1 <= n <= 45
//
//
// Related Topics 记忆化搜索 数学 动态规划 👍 2895 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {

    /**
    * @param Integer $n
    * @return Integer
    */
    function climbStairs($n) {
        /*方案一：
        if($n<0) return 0;
        return $n <= 1 ? 1 : ($this->climbStairs($n-1)+$this->climbStairs($n-2));*/

        
        $q = $r = 0;
        $p = 1;
        for($i=0;$i<$n;$i++){
            $r = $q;
            $q = $p;
            $p = $q + $r;
        }
        return $p;
    
    }
}
//leetcode submit region end(Prohibit modification and deletion)

$a = new Solution();
$level = 10;
var_dump($a->climbStairs($level));

/*
解题思路：
1、之前面试做过，斐波那契数列变体，使用递归来做，两行代码可以直接写完，边界值的话，小于0返回0，等于0返回1，（当$n等于44的时候，本地测试需要1m7.311s，该方法废弃，但是面试时仍旧可以当作一个方案来说）
2、参考官方的解题思路，
笨办法，手工先进行几步：
1阶阶梯只有一种方案
2阶则有2种， 1+2
3阶则有3种 1+1+1 1+2 2+1
4阶则有5种 1+1+1+1 2+1+1 1+2+1 1+1+2 2+2
5阶则有8种 1+1+1+1+1 2+1+1+1 1+2+1+1 1+1+2+1 2+2+1  1+1+1+2 1+2+2 2+1+2
....
列出公式则为
f(x) = f(x-1)+f(x-2)
对于上面公式的理解：
如果你仔细看我的第3阶和第4阶以及第5阶，你是否会发现 第5阶的前5种方案都是在第4阶的基础之上再增加一步，后3种方案是在第3阶上又增加了两步，因为每次的选择只能前进一步或者两步，
所以，该方法不会导致方案重复。
对于上面的公式我们来进行转化，
p=q+r
即最后的值结果为p，f(x-1)为q，f(x-2)为r，每次结果返回s
复制leetcode官方的图片如下：
https://assets.leetcode-cn.com/solution-static/70/70_fig1.gif

*/
