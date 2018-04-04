package doublelink

type Element struct {
	prev, next *Element
	Value interface{}
}

func NewElement(val interface{}) (*Element) {
	return &Element{nil, nil, val}
}

func (e *Element) traverseR() (*Element) {
	loc := e
	for loc.next != nil {
		loc = loc.next
	}
	return loc
}

func (e *Element) traverseL() (*Element) {
	loc := e
	for loc.prev != nil {
		loc = loc.prev
	}
	return loc
}

func (e *Element) LeftRemove() (*Element, interface{}) {
	if e != nil {
		loc := e.traverseL()
		next := loc.next
		if next != nil {
			next.prev = nil
		}
		return next, loc.Value
	}
	return nil, nil
}

func (e *Element) LeftPut(val interface{}) (*Element) {
	if e != nil {
		loc := e.traverseL()
		tmp := &Element{nil, loc, val}
		loc.prev = tmp
		return tmp
	} else {
		tmp := &Element{nil, nil, val}
		return tmp
	}
}

func (e *Element) RightPut(val interface{}) (*Element) {
	if e != nil {
		loc := e.traverseR()
		tmp := &Element{loc, nil, val}
		loc.next = tmp
		return tmp
	} else {
		tmp := &Element{nil, nil, val}
		return tmp
	}
}

func (e *Element) RightRemove() (*Element, interface{}) {
	if e != nil {
		loc := e.traverseR()
		prev := loc.prev
		if prev != nil {
			prev.next = nil
		}
		return prev, loc.Value
	}
	return nil, nil
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
