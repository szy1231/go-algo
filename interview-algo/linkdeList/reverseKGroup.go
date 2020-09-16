package linkdeList

//K 个一组翻转链表

/**
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。



示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5



说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

//无头链表反转
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	begin := head
	var end *ListNode
	var next *ListNode
	pre := &ListNode{Next: head}
	res := pre

	for begin != nil {

		end = begin
		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				pre.Next = begin
				return res.Next
			}
		}
		next = end.Next
		end.Next = nil
		//无头链表反转
		pre.Next = reverse(begin)
		pre = begin
		begin = next
	}

	return res.Next
}
