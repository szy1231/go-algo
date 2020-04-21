/*
@Time : 2020/4/21 17:29
@Author : pasiyu
@File : bubbleSort_test.go
*/
package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	a := []int{9, 4, 7, 0, 6, 8, 4}
	BubbleSort(a)
	t.Log(a)
}
