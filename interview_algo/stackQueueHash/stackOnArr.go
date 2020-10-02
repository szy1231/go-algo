package stackQueueHash

import "errors"

type SliceStack struct {
	arr       []int
	stackSize int
}

func (s *SliceStack) IsEmpty() bool {
	return s.stackSize == 0
}

func (s *SliceStack) Size() int {
	return s.stackSize
}

func (s *SliceStack) Top() int {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	return s.arr[s.stackSize-1]
}

func (s *SliceStack) Pop() int {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	s.stackSize--
	tmp := s.arr[s.stackSize]
	s.arr = s.arr[:s.stackSize]
	return tmp
}

func (s *SliceStack) Push(t int) {
	s.arr = append(s.arr, t)
	s.stackSize = s.stackSize + 1
}
