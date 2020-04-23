/*
@Time : 2020/4/23 16:46
@Author : pasiyu
@File : quickSort_test.go
*/
package sorts

import "testing"

func TestQuickSort(t *testing.T) {
	a := []int{9, 4, 7, 0, 6, 8, 4}
	QuickSort(a)
	t.Log(a)
}
