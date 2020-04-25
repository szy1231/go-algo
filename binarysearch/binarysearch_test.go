/*
@Time : 2020/4/25 9:46
@Author : pasiyu
@File : binarysearch_test.go
*/
package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1,2,5,6,7,8,9,10}
	t.Log(BinarySearch(a, 1))
}

func TestBinarySearchCycle(t *testing.T) {
	a := []int{1,2,5,6,7,8,9,10}
	t.Log(BinarySearchCycle(a, 3))
}
