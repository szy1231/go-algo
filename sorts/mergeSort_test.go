/*
@Time : 2020/4/23 11:45
@Author : pasiyu
@File : mergeSort_test.go
*/
package sorts

import "testing"

func TestMergeSort(t *testing.T) {
	a := []int{9, 4, 7, 0, 6, 8, 4}
	MergeSort(a)
	t.Log(a)
}
