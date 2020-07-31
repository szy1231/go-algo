package offer

/*
//剑指 Offer 04. 二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。



限制：

0 <= n <= 1000

0 <= m <= 1000

*/
func findNumberIn2DArray(matrix [][]int, target int) bool {

	//m代表第一层切片的长度
	m := len(matrix)
	if m < 1 {
		return false
	}
	//n代表第二层切片的长度
	n := len(matrix[0])
	//最后一行最后一列的元素都比他小，说明没有满足条件的元素
	if n < 1 || matrix[m-1][n-1] < target {
		return false
	}

	//遍历切片
	for i := 0; i < m; i++ {
		//如果这一行的第一个元素都比他大的话，后面肯定没有满足的元素
		if target < matrix[i][0] {
			return false
		}
		for j := 0; j < n; j++ {
			//如果比这个切片第一个还小，或者比最后一个还大。说明这个切片中没有满足条件的，直接break
			if target < matrix[i][j] || target > matrix[i][n-1] {
				break
			}
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false

}
