package leetcode

import (
	"container/list"
)

//leetcode102题

/*
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//利用队列先进先出的特点，把二叉树依次存入队列
func levelOrder(root *TreeNode) [][]int {
	//创建一个保存最终结果的二位切片
	var result [][]int

	//如果节点为空，直接返回
	if root == nil {
		return result
	}

	//创建一个队列，并把根节点从前端入队
	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		//保存当前层数据的切片
		var curLevel []int
		//count：保存当前层节点个数
		count := queue.Len()

		for count > 0 {
			//从队列后端取出数据
			element := queue.Back()
			node := element.Value.(*TreeNode)
			curLevel = append(curLevel, node.Val)

			//把左右子节点依次从队列前端入队
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}

			//删除节点，并把当前剩余节点数-1
			queue.Remove(element)
			count--
		}

		result = append(result, curLevel)
	}
	return result
}
