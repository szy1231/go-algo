/*
@Time : 2020/4/10 9:28
@Author : pasiyu
@File : singlelinkedlist.go
*/
package main

import "fmt"

/*
 单链表基本操作
*/

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}
//创建一个新节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}
//创建一个新链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}
//获取值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	newNode.next = p.next
	p.next = newNode
	this.length++
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//链表反转
func (this *LinkedList) Reversal() {
	if this == nil || this.head.next == nil {
		return
	}

	var preNode *ListNode //之前
	var cur = this.head.next
	var nextNode *ListNode = nil

	for cur != nil {
		nextNode = cur.next
		cur.next = preNode
		preNode = cur
		cur = nextNode
	}
	this.head.next.next = nil
	this.head.next = preNode
}
