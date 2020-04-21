/*
@Time : 2020/4/21 22:27
@Author : pasiyu
@File : insertSort_test.go
*/
package sorts

import "testing"

func TestInsertSort(t *testing.T) {
	a := []int{9, 4, 7, 0, 6, 8, 4}
	InsertSort(a)
	t.Log(a)
}
