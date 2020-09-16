package leetcode

//21. 合并两个有序链表
/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

//遍历的同时比较l1与l2的大小
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	cur := new(ListNode)
	res := cur

	for l1 != nil && l2 != nil {
		//比较l1和l2的大小
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	//看l1是否有剩余节点
	if l1 != nil {
		cur.Next = l1
	}
	//看l2是否有剩余节点
	if l2 != nil {
		cur.Next = l2
	}

	return res.Next
}
