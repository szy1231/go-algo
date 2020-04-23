/*
@Time : 2020/4/22 16:43
@Author : pasiyu
@File : selectSort_test.go
*/
package sorts

import "testing"

func TestSelectSort(t *testing.T) {
	a := []int{9, 4, 7, 0, 6, 8, 4}
	SelectSort(a)
	t.Log(a)
}
