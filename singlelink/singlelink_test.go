package singlelink

import "testing"

func TestNewElement(t *testing.T) {
	item := "Hi"
	n := NewElement(item)
	if n.value != item {
		t.Error("where did the item to store go?")
	}
	if n.next != nil {
		t.Error("a singleton node shouldn't have a next")
	}
}
