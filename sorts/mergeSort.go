/*
@Time : 2020/4/23 11:23
@Author : pasiyu
@File : mergeSort.go
*/
package sorts

//简单过程
/*
	11 8 3 9 7 1 2 5
  11 8 3 9    7 1 2 5
 11 8  3 9  7 1  2 5
 8 11  3 9  1 7  2 5
  3 8 9 11    1 2 5 7
    1 2 3 5 7 8 9 11
*/

/*
  归并排序:如果要排序一个数组，我们先把数组从中间分成前后
  两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起
*/
func MergeSort(a []int) {
	n := len(a)

	if n <= 1 {
		return
	}
	mergeSort(a, 0, n-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)

	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	//排序
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if a[i] <= a[j] {
			tmpArr[k] = a[i]
			i++
		} else {
			tmpArr[k] = a[j]
			j++
		}
	}

	//检查是否有剩余元素
	for ; i <= mid; i++ {
		tmpArr[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		tmpArr[k] = a[j]
		k++
	}
	//从tmpArr拷贝到a元切片中
	copy(a[start:end+1], tmpArr)
}
