package singlelink

type Element struct {
	value interface{}
	next *Element
}

func NewElement(val interface{}) *Element{
	return &Element{val, nil}
}

func (l *Element) Add(val interface{}) *Element{
	elt := NewElement(val)
	elt.next = l
	return elt
}

func (l *Element) Remove() (val interface{}, next *Element) {
	val = l.value
	next = l.next
	return
}

func (l *Element) CutNext() {
	if l.next != nil {
		l.next = l.next.next
	}
}

func (l *Element) Forward(dist int) *Element{
	cur := l
	for (dist > 0 && cur.next != nil) {
		cur = cur.next
		dist--
	}
	return cur
}

