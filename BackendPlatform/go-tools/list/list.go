package list

import "container/list"

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Offer(v interface{}) {
	q.list.PushBack(v)
}
func (q *Queue) Poll() interface{} {
	v := q.list.Front()
	if v != nil {
		q.list.Remove(v)
		return v.Value
	}
	return nil

}
func (q *Queue) Len() int {
	return q.list.Len()
}
func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	e := s.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
