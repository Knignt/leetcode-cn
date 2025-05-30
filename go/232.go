package main

import "fmt"

/*
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


示例 1：

输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false


提示：

1 <= x <= 9
最多调用 100 次 push、pop、peek 和 empty
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）


进阶：

你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
*/

type MyQueue struct {
	in  *Stack //数据存储
	out *Stack //输出结果
}

type Stack struct {
	data []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{
		in: &Stack{
			make([]int, 0),
		},
		out: &Stack{
			make([]int, 0),
		},
	}
}

func (this *MyQueue) Push(x int) {
	this.in.data = append(this.in.data, x)
}

func (this *MyQueue) transfer() {

	if len(this.out.data) == 0 { //只有out全部pop完毕，才会再从in里面在进行数据转移
		n := len(this.in.data) - 1
		for n >= 0 {
			this.out.data = append(this.out.data, this.in.data[n])
			n--
		}

		this.in.data = make([]int, 0) //因为in里面的数据全部转移到了out里面,所以这里in给到了0
	}
}

func (this *MyQueue) Pop() int {
	this.transfer()
	tmp := this.out.data[len(this.out.data)-1]
	this.out.data = this.out.data[0 : len(this.out.data)-1]
	return tmp
}

func (this *MyQueue) Peek() int {
	this.transfer()
	tmp := this.out.data[len(this.out.data)-1]
	return tmp
}

func (this *MyQueue) Empty() bool {
	return len(this.in.data) == 0 && len(this.out.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main232() {
	obj := ConstructorQueue()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)

	fmt.Println(obj.Pop())

	obj.Push(5)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
}
