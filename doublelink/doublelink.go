package doublelink

type Element struct {
	prev, next *Element
	Value interface{}
}

func NewElement(val interface{}) *Element {
	return &Element{nil, nil, val}
}

func (e *Element) Prepend(val interface{}) (*Element) {
	loc := e
	for loc.prev != nil {
		loc = loc.prev
	}
	tmp := &Element{nil, loc, val}
	loc.prev = tmp
	return tmp
}

func (e *Element) Postpend(val interface{}) (*Element) {
	loc := e
	for loc.next != nil {
		loc = loc.next
	}
	tmp := &Element{loc, nil, val}
	loc.next = tmp
	return tmp
}

func (e *Element) Forward(dist int) (*Element){
	tmp := e
	for (tmp.next != nil && dist > 0) {
		tmp = tmp.next
		dist--
	}
	return tmp
}

func (e *Element) Backward(dist int) (*Element){
	tmp := e
	for (tmp.prev != nil && dist > 0) {
		tmp = tmp.prev
		dist--
	}
	return tmp
}
