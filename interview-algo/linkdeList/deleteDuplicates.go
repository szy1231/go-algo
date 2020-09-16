package linkdeList

//删除无序链表中的重复元素
//分析：1：循环暴力删除 时间复杂度O(n2)
//		2:递归	时间复杂度O(n2)
//  	3：利用集合去重 时间复杂度O(n)
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	hashMap := make(map[int]struct{})

	res := head
	var pre *ListNode
	for head != nil {
		if _, ok := hashMap[head.Val]; ok {
			pre.Next = head.Next
			head = head.Next
			continue
		}
		hashMap[head.Val] = struct{}{}
		pre = head
		head = head.Next
	}

	return res
}

//删除有序列表中的重复元素
//时间复杂度O(n)
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	next := head.Next
	for next != nil {
		if next.Val != cur.Val {
			cur.Next = next
			cur = next
			next = next.Next
			continue
		}
		next = next.Next
	}
	cur.Next = nil
	return head
}
