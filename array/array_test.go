/*
@Time : 2020/4/9 16:43
@Author : pasiyu
@File : array_test.go
*/
package array

import (
	"testing"
)

func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := NewArray(int(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(int(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(int(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := NewArray(int(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(int(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Find(2))
	t.Log(arr.Find(5))
	t.Log(arr.Find(15))
}

func BenchmarkArray_Find(b *testing.B) {
	capacity := 10
	arr := NewArray(int(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(int(i), i+1)
		if nil != err {
			b.Fatal(err.Error())
		}
	}
	for i := 0; i < b.N; i++ {
		_, err := arr.Find(9)
		if err != nil {
			b.Error(err.Error())
		}
	}
}
