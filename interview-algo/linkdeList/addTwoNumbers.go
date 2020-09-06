package linkdeList

//链表求和
/**
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912

*/

//思路
/**
1：整数相加 遍历链表，求出链表代表的整数，然后相加，最后用链表形式返回
2：链表相加：直接两个链表相加 生成新的链表
*/

//链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l2
	}

	//记录进位
	c := 0
	//记录两链表的和
	sum := 0

	//最后结果的头节点
	resHead := &ListNode{}
	p := resHead
	for l1 != nil && l2 != nil {
		p.Next = &ListNode{}
		p = p.Next
		sum = l1.Val + l2.Val + c
		c = sum / 10
		p.Val = sum % 10

		l1 = l1.Next
		l2 = l2.Next
	}

	//查看l1是否有剩余节点
	if l1 != nil {
		for l1 != nil {
			p.Next = &ListNode{}
			p = p.Next
			sum = l1.Val + c
			c = sum / 10
			p.Val = sum % 10

			l1 = l1.Next
		}
	}
	//查看l2是否有剩余节点
	if l2 != nil {
		for l2 != nil {
			p.Next = &ListNode{}
			p = p.Next
			sum = l2.Val + c
			c = sum / 10
			p.Val = sum % 10

			l2 = l2.Next
		}
	}

	//查看最后进位是否为空
	if c != 0 {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = c
	}

	//第一个节点为空节点，返回next
	return resHead.Next
}