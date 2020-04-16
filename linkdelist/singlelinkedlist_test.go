/*
@Time : 2020/4/15 17:16
@Author : pasiyu
@File : singlelinkedlist_test.go
*/
package main

import "testing"

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func BenchmarkLinkedList_InsertToTail(b *testing.B) {
	l := NewLinkedList()
	for i := 0; i < b.N; i++ {
		l.InsertToTail(i + 1)
	}
}

func TestReversal(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	l.Reversal()
	l.Print()
}
