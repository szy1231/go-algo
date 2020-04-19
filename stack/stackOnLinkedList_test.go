/*
@Time : 2020/4/19 18:40
@Author : pasiyu
@File : stackOnLinkedList_test.go
*/
package stack

import "testing"

func TestLinkedListStack_Print(t *testing.T) {
	s := NewLinkedListStack()
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
