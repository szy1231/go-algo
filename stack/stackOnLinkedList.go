/*
@Time : 2020/4/19 18:18
@Author : pasiyu
@File : stackOnLinkedList.go
*/
package stack

import "fmt"

/*
基于链表实现的数组
*/

type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	topNode *node //	栈顶节点
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) IsEmpty() bool {
	return this.topNode == nil
}

func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &node{this.topNode, v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode.val
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("空栈")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
