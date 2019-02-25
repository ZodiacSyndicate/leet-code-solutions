/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (47.35%)
 * Total Accepted:    15.7K
 * Total Submissions: 33.2K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 * 
 * 
 * 示例:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 */
import "math"

type MinStack struct {
	Val []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack { Val: []int{}}
}


func (this *MinStack) Push(x int)  {
  this.Val = append(this.Val, x)
}


func (this *MinStack) Pop()  {
  this.Val = this.Val[:len(this.Val)-1]
}


func (this *MinStack) Top() int {
  return this.Val[len(this.Val)-1]
}


func (this *MinStack) GetMin() int {
	res := float64(math.Pow(2, 31))
  for _, n := range this.Val {
		res = math.Min(float64(n), res)
	}
	return int(res)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

