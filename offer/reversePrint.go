package offer

/*
剑指 Offer 06
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

限制：

0 <= 链表长度 <= 10000
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//解法一：1、遍历链表，同时加入数组
//		2、反转数组
func reversePrint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res

}

//解法2：遍历数组加入堆中，从堆中取出加入数组
type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}
func (s *stack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) len() int {
	return len(*s)
}
func reversePrint01(head *ListNode) []int {
	stack := new(stack)
	for head != nil {
		stack.push(head.Val)
		head = head.Next
	}

	l := stack.len()
	res := make([]int, l)

	for i := 0; i < l; i++ {
		res[i] = stack.pop()
	}
	return res

}
