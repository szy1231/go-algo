/*
@Time : 2020/4/21 17:09
@Author : pasiyu
@File : bubbleSort.go
*/
package sorts

//简单过程
/*
	4 5 6 3 2 1
	4 5 3 2 1 6
    4 3 2 1 5 6
	3 2 1 4 5 6
 	2 1 3 4 5 6
	1 2 3 4 5 6
*/
//冒泡排序
func BubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		//如果一轮没有交换,则退出循环
		flag := false
		for j := 0; j < n-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
