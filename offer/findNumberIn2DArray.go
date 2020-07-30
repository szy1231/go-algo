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

	//如果是空，则没有查找的意义，直接返回
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	//确定行的范围
	n := 0
	i := 0
	for ; n < len(matrix); n++ {
		if matrix[n][0] < target {
			if matrix[n][len(matrix[0])-1] < target {
				i++
			}
			continue
		}
		if matrix[n][0] == target {
			return true
		}
	}

	//确定列的范围
	l := 0
	for ; l < len(matrix[0]); l++ {
		if matrix[0][l] < target {
			continue
		}
		if matrix[0][l] == target {
			return true
		}
	}

	//在列和行中查找
	for ; i < n; i++ {
		for j := 0; j < l; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false

}
