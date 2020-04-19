/*
@Time : 2020/4/19 16:32
@Author : pasiyu
@File : stackOnArray.go
*/
package stack

import "fmt"

/*
基于数组实现一个栈
*/

type ArrayStack struct {
	data []interface{} //数据
	top  int           //栈顶指针
}

//初始化一个栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//是否为空栈
func (this *ArrayStack) IsEmpty() bool {
	return this.top < 0
}

//入栈
func (this *ArrayStack) Push(v interface{}) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top += 1
	}

	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

//出栈
func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top -= 1
	return v
}

//查看栈顶数据
func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

//清空栈
func (this *ArrayStack) Flush() {
	this.top = -1
}

//打印
func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("空栈")
	} else {
		for i := this.top; i > -1; i-- {
			fmt.Println(this.data[i])
		}
	}
}
