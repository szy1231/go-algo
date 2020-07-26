package leetcode

import "container/list"

//104. 二叉树的最大深度
/*
*给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3

   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

*/
// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//1：递归解法
func maxDepth01(root *TreeNode) int {
	return recursive(root, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func recursive(root *TreeNode, n int) int {
	if root == nil {
		return n - 1
	}

	if root.Left == nil && root.Right == nil {
		return n
	}

	return max(recursive(root.Left, n+1), recursive(root.Right, n+1))

}

//2:BFS
func maxDepth02(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushFront(root)
	//记录深度
	depth := 0

	for queue.Len() > 0 {
		//当前层的数量
		count := queue.Len()
		for count > 0 {

			element := queue.Back()
			queue.Remove(element)
			node := element.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}

			if node.Right != nil {
				queue.PushFront(node.Right)
			}

			count--
		}
		depth++
	}
	return depth

}
