package doublelink

import "errors"

type Queue struct {
	head, tail *Element
}

func NewQueue() (*Queue) {
	return &Queue{nil, nil}
}

func (q *Queue) Empty() bool {
	return q.head == nil && q.tail == nil
}

func (q *Queue) Put(item interface{}) {
	q.tail = q.tail.RightPut(item)
	if q.head == nil {
		q.head = q.tail
	}
}

func (q *Queue) Get() (item interface{}, err error) {
	if q.head == nil {
		return nil, errors.New("Get from empty queue")
	} else {
		q.head, item = q.head.LeftRemove()
		if q.head == nil {
			q.tail = nil
		}
		return item, nil
	}
}
