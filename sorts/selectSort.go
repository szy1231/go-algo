/*
@Time : 2020/4/22 16:34
@Author : pasiyu
@File : selectSort.go
*/
package sorts

//选择排序
func SelectSort(a []int) {
	n := len(a)

	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		//记录最小值的索引
		min := i
		for j := i + 1; j < n; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
