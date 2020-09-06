package linkdeList

//重排链表
/**
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

*/

/**
解题思路：找到链表的中间节点，对后半部分链表进行逆序。然后把前后半部分节点进行合并
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := findMidNode(head)
	cur1 := head
	cur2 := reverse(mid)

	var tmp *ListNode

	//合并
	for cur1.Next != nil {
		tmp = cur1.Next
		cur1.Next = cur2
		cur1 = tmp
		tmp = cur2.Next
		cur2.Next = cur1
		cur2 = tmp
	}
	cur1.Next = cur2

	return
}

//寻找中间节点
func findMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//快慢指针
	slow, fast, slowPre := head, head, head

	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowPre.Next = nil
	return slow
}

//反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
