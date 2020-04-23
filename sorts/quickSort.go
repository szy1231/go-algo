/*
@Time : 2020/4/23 15:18
@Author : pasiyu
@File : quickSort.go
*/
package sorts

/*快速排序:如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r
  之间的任意一个数据作为 pivot（分区点）。*/
func QuickSort(a []int) {
	separateSort(a, 0, len(a)-1)
}

func separateSort(a []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(a, start, end)
	separateSort(a, 0, i-1)
	separateSort(a, i+1, end)

}

func partition(a []int, start, end int) int {
	//选取最后一位数字当对比
	pivot := a[end]

	i := start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
			i++
		}
	}
	//把对比的数放中间
	a[i], a[end] = a[end], a[i]

	return i
}
