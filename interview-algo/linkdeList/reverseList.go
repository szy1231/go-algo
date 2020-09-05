package linkdeList

//链表反转（无头节点）
/**
7 6 5 4 3 2 1  ---->  1 2 3 4 5 6 7
*/


type ListNode struct {
	Val  int
	Next *ListNode
}

//1：就地反转，保存前驱节点和后继节点
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	var cur *ListNode

	next := head

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}

	return pre
}

//2：递归
/**
1 2 3 4 5
1 5 4 3 2
5 4 3 2 1
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := RecursiveReverseChild(head)
	return newHead
}

func RecursiveReverseChild(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := RecursiveReverseChild(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

//3:插入法
/**
1 2 3 4 5
2 1 3 4 5
3 2 1 4 5
4 3 2 1 5
5 4 3 2 1
*/
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var cur *ListNode
	var next *ListNode
	cur = head.Next
	head.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = head
		head = cur
		cur = next
	}

	return head
}
