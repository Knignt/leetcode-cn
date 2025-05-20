package main

import (
	"container/list"
	"fmt"
)

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/
type LRUCache struct {
	ll       *list.List
	hash     map[int]*list.Element
	capacity int
}

type LRUCacheItem struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ll:       list.New(),
		hash:     make(map[int]*list.Element),
		capacity: capacity,
	}

}

func (this *LRUCache) Get(key int) int {
	if elem, exist := this.hash[key]; exist {
		this.ll.MoveToFront(elem)
		this.hash[key] = this.ll.Front()
		return this.hash[key].Value.(*LRUCacheItem).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if ele, exists := this.hash[key]; exists {
		//值的更改
		ele.Value.(*LRUCacheItem).value = value
		this.ll.MoveToFront(ele)
		this.hash[key] = this.ll.Front()
		return
	}

	item := LRUCacheItem{
		key:   key,
		value: value,
	}

	elem := this.ll.PushFront(&item)
	this.hash[key] = elem

	if this.capacity < len(this.hash) {
		//置换
		tmp := this.ll.Back()
		this.ll.Remove(this.ll.Back())
		delete(this.hash, tmp.Value.(*LRUCacheItem).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)

	lRUCache.Put(2, 1)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 3)
	lRUCache.Put(4, 1)
	fmt.Println(lRUCache.Get(1)) //-1
	fmt.Println(lRUCache.Get(2)) //3

	/*lRUCache.Put(1, 1)           // 缓存是 {1=1}
	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4*/
}

func printContainerList1(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\t", e.Value)
	}
	fmt.Println()
}
