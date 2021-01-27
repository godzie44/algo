package linkedlist

import "errors"

type dllNode struct {
	key  int
	prev *dllNode
	next *dllNode
}

type DoubleLinkedList struct {
	Head *dllNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) Insert(x *dllNode) {
	if d.Head != nil {
		x.next = d.Head
		d.Head.prev = x
	}

	x.prev = nil
	d.Head = x
}

var ErrNodeNotFound = errors.New("node not found")

func (d *DoubleLinkedList) Search(key int) (*dllNode, error) {
	for iter := d.Head; iter != nil; iter = iter.next {
		if iter.key == key {
			return iter, nil
		}
	}

	return nil, ErrNodeNotFound
}

func (d *DoubleLinkedList) Delete(x *dllNode) {
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		d.Head = x.next
		x.prev = nil
	}

	if x.next != nil {
		x.next.prev = x.prev
	}
}

func (d *DoubleLinkedList) IntoArray() (arr []int) {
	for iter := d.Head; iter != nil; iter = iter.next {
		arr = append(arr, iter.key)
	}
	return
}

func (d *DoubleLinkedList) Revert() {
	iter := d.Head
	for iter.next != nil {
		x := iter
		iter = iter.next

		x.next, x.prev = x.prev, x.next
	}

	iter.next, iter.prev = iter.prev, iter.next
	d.Head = iter
}
