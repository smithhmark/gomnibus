package doublelink

import "testing"

func expectNoError(exp, rcv interface{}, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Should not have an error here")
	}
	if exp != rcv {
		t.Fatalf("Expected (%v) != received(%v)", exp, rcv)
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Put("Micky")
	q.Put("Minny")
	q.Put("Daffy")

	{
		v,e := q.Get(); 
		expectNoError(v, "Micky", e, t)
	}
	{
		v,e := q.Get(); 
		expectNoError(v, "Minny", e, t)
	}
	{
		v,e := q.Get(); 
		expectNoError(v, "Daffy", e, t)
	}
	if q.Empty() != true {
		t.Fatalf("Q.head:%v Q.tail:%v", q.head, q.tail)
	}
}
