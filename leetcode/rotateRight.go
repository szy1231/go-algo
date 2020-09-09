package leetcode

//61. 旋转链表

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

//[0,1,2]
//4
//201
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow,fast := head,head
	for ;k>0 ;k-- {
		if fast.Next == nil {
			fast = head
			continue
		}
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Next == nil {
		return head
	}

	res := slow.Next
	fast.Next = head
	slow.Next = nil
	return res
}
