package singlelink

import "errors"

type Stack struct {
	head   *Element
	length int
}

func NewStack() *Stack {
	s := &Stack{nil, 0}
	return s
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Push(item interface{}) {
	newHead := s.head.Add(item)
	s.head, s.length = newHead, s.length+1
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return nil, errors.New("Pop on empty Stack")
	}
	val, newhead := s.head.Remove()
	s.head = newhead
	s.length--
	return val, nil
}
