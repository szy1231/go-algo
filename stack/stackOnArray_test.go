/*
@Time : 2020/4/19 16:52
@Author : pasiyu
@File : stackOnArray_test.go
*/
package stack

import "testing"

func TestArrayStack_Print(t *testing.T) {
	s := NewArrayStack()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Print()
	t.Log(s.Top())
	t.Log(s.Pop())
	s.Print()
	s.Flush()
	s.Print()
	t.Log(s.Top())
}