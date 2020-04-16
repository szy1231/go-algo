/*
@Time : 2020/4/9 12:47
@Author : pasiyu
@File : array.go
*/
package array

import (
	"errors"
	"fmt"
)

/*
 数组的基本操作
 */

type Array struct {
	data   []int
	length int
}

//为数组初始化内存
func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (this *Array) Len() int {
	return this.length
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index int) bool {
	if index >= int(cap(this.data)) {
		return true
	}
	return false
}

//通过索引查找数组
func (this *Array) Find(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("索引越界")
	}
	return this.data[index], nil
}

//插入数值到索引index上
func (this *Array) Insert(index int, v int) error {
	if this.Len() == int(cap(this.data)) {
		return errors.New("数组已满")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("索引越界")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

//最后添加
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

//删除索引上的值
func (this *Array) Delete(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("索引越界")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

//打印数列
func (this *Array) Print() {
	var format string
	for i := int(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%v", this.data[i])
	}
	fmt.Println(format)
}
