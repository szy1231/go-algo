/*
@Time : 2020/4/20 13:40
@Author : pasiyu
@File : queueOnArray.go
*/
package queue


type ArrayQueue struct {
	q        []interface{} //数据
	capacity int           //容量
	head     int           //出队列的下标
	tail     int           //进队列的下标
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (this *ArrayQueue) Enqueue(v interface{}) bool {
	if this.tail == this.capacity {
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) Dequeue() interface{} {
	if this.head == this.tail {
		return nil
	}
	v := this.q[this.head]
	this.head++
	return v
}
