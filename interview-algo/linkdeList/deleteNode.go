package linkdeList

//删除链表中的节点

//把下一个节点的信息复制到本节点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
