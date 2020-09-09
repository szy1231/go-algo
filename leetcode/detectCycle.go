package leetcode

//142. 环形链表 II
/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。
*/

//1：利用map存储遍历过的节点，如果在遍历的同时在map取出相同的节点，第一个取出的节点则为第一个节点
//空间复杂度为O(n)
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	hashMap := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return head
		}
		hashMap[head] = struct{}{}
		head = head.Next
	}

	return nil
}

//快慢指针
func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	for head != fast {
		head = head.Next
		fast = fast.Next
	}

	return head
}
