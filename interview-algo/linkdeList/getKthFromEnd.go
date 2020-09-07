package linkdeList

//链表中倒数第k个节点

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。



示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/

//解题思路
/**
第一种：遍历链表，得出链表长度n。利用n-k取得倒数第k元素。需要遍历2次
第二种：前后指针，相差k，当后指针到达末尾时，返回前指针。效率更高
*/

//第二种方案
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	bef := head
	for ; k > 0 ; k-- {
		head = head.Next
	}

	for head != nil {
		bef = bef.Next
		head = head.Next
	}

	return bef
}
