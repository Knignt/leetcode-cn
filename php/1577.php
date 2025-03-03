<?php

/*
给你一个 有向无环图 ， n 个节点编号为 0 到 n-1 ，以及一个边数组 edges ，其中 edges[i] = [fromi, toi] 表示一条从点  fromi 到点 toi 的有向边。

找到最小的点集使得从这些点出发能到达图中所有点。题目保证解存在且唯一。

你可以以任意顺序返回这些节点编号。

 

示例 1：



输入：n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
输出：[0,3]
解释：从单个节点出发无法到达所有节点。从 0 出发我们可以到达 [0,1,2,5] 。从 3 出发我们可以到达 [3,4,2,5] 。所以我们输出 [0,3] 。
示例 2：



输入：n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]
输出：[0,2,3]
解释：注意到节点 0，3 和 2 无法从其他节点到达，所以我们必须将它们包含在结果点集中，这些点都能到达节点 1 和 4 。
 

提示：

2 <= n <= 10^5
1 <= edges.length <= min(10^5, n * (n - 1) / 2)
edges[i].length == 2
0 <= fromi, toi < n
所有点对 (fromi, toi) 互不相同。
*/

class Solution {

    /**
     * @param Integer $n
     * @param Integer[][] $edges
     * @return Integer[]
     */
    function findSmallestSetOfVertices($n, $edges) {
        if(empty($edges)) return [];

        $nodes = [];
        foreach($edges as $row){
            if(!isset($nodes[$row[1]])) $nodes[$row[1]] = 0;
            $nodes[$row[1]]++;
        }

        $res = [];
        for($i=0;$i<$n;$i++){
            if(!isset($nodes[$i])){
                $res[] = $i;
            }
        }

        return $res;
    }
}

$a = new Solution();
$n = 6;
$edges = [[0,1],[0,2],[2,5],[3,4],[4,2]];
var_dump($a->findSmallestSetOfVertices($n, $edges));

//中等难度的题目
/*
刚开始的思路是，将每个节点的入口和出口，分别弄一个数组，只要入口的不在出口的，返回即可，btw，最后题解时没想到居然有一个100000大小的数组，直接给干的超出时间限制
没得办法，去看题解，题解很简单，只要有出口的，放到一个数组里面，进行加加，剩下的就是入口了，因为 提前给了节点的范围，（即第一个值$n）
*/
