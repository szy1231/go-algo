/*
@Time : 2020/4/21 22:12
@Author : pasiyu
@File : insertSort.go
*/
package sorts

func InsertSort(a []int) {
	n := len(a)

	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if v > a[j] {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
}
