package singlelink

//import "fmt"
import "errors"

type Queue interface {
	Put(interface{})
	Get() (interface {}, error)
	Size() int
}

type QueueFactory func () Queue

func rev(src, dest *Stack) {
	for src.Size() > 0 {
		val, _ := src.Pop()
		dest.Push(val)
	}
}

type ModeQueue struct {
	a *Stack
	b *Stack
	putMode bool
}

func NewModeQueue() Queue {
	q := &ModeQueue{NewStack(), NewStack(), true}
	return q
}

func (q *ModeQueue) Size() int {
	if q.putMode {
		return q.a.Size()
	} else {
		return q.b.Size()
	}
}

func (q *ModeQueue) Put(item interface{}) {
	if !q.putMode {
		q.putMode = true
		rev(q.b, q.a)
	}
	q.a.Push(item)
}

func (q *ModeQueue) Get() (ret interface{}, err error) {
	if q.putMode {
		rev(q.a, q.b)
		q.putMode = false
	}
	if q.b.Size() == 0 {
		return 0, errors.New("Queue is empty")
	}

	ret, err = q.b.Pop()
	if err != nil {
		return 0, errors.New("Something smells in Denmark...")
	}
	err = nil
	return
}
