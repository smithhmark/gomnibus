package doublelink

import "testing"

func TestConstruction(t *testing.T) {
	val := "Mario"
	e := NewElement(val)

	if e.next != nil {
		t.Fatalf("next should be nil")
	}
	if e.prev != nil {
		t.Fatalf("prev should be nil")
	}
	if e.Value != val {
		t.Fatalf("lost the stored value")
	}
}

func TestPrepend(t *testing.T) {
	val := []string{
		"Mario",
		"Luigi",
		"Peach",
	}
	e := NewElement(val[0])

	luigi := e.Prepend(val[1])
	if luigi.Value != val[1] || luigi.prev != nil || luigi.next != e {
		t.Fatalf("should have created a {nil, <mario>, %s} element", val[1])
	}

	peach := e.Prepend(val[2])
	if peach.Value != val[2] || peach.prev != nil || peach.next != luigi {
		t.Fatalf("should have created a {nil, <luigi>, %s} element", val[2])
	}
}

func TestPostpend(t *testing.T) {
	val := []string{
		"Mario",
		"Luigi",
		"Peach",
	}
	e := NewElement(val[0])

	luigi := e.Postpend(val[1])
	if luigi.Value != val[1] || luigi.next != nil || luigi.prev != e {
		t.Fatalf("should have created a {nil, <mario>, %s} element", val[1])
	}

	peach := e.Postpend(val[2])
	if peach.Value != val[2] || peach.next != nil || peach.prev != luigi {
		t.Fatalf("should have created a {nil, <luigi>, %s} element", val[2])
	}
}

func TestForward(t *testing.T) {
	m := NewElement("Mario")
	head := m
	tail := head
	l := tail.Postpend("Luigi")
	tail = l
	p := head.Prepend("Peach")
	head = p

	if head.Forward(1) != m {
		t.Fatalf("Forward(1) should have taken us to <Mario>")
	}
	if head.Forward(2) != l {
		t.Fatalf("Forward(2) should have taken us to <Luigi>")
	}
	if head.Forward(3) != l {
		t.Fatalf("Forward(3) should have taken us to <Luigi>")
	}
}

func TestBackward(t *testing.T) {
	m := NewElement("Mario")
	head := m
	tail := head
	l := tail.Postpend("Luigi")
	tail = l
	p := head.Prepend("Peach")
	head = p

	if tail.Backward(1) != m {
		t.Fatalf("Backward(1) should have taken us to <Mario>")
	}
	if tail.Backward(2) != p {
		t.Fatalf("Backward(2) should have taken us to <Peach>")
	}
	if tail.Backward(3) != p {
		t.Fatalf("Backward(3) should have taken us to <Peach>")
	}
}
