package linkedlist

type llNode struct {
	next *llNode
	key  int
}

type LinkedList struct {
	Head *llNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (d *LinkedList) Insert(x *llNode) {
	x.next = d.Head
	d.Head = x
}

func (d *LinkedList) Search(key int) (*llNode, error) {
	for iter := d.Head; iter != nil; iter = iter.next {
		if iter.key == key {
			return iter, nil
		}
	}

	return nil, ErrNodeNotFound
}

func (d *LinkedList) Delete(x *llNode) {
	var p **llNode
	p = &d.Head

	for *p != nil {
		if *p == x {
			*p = x.next
			break
		}
		p = &(*p).next
	}
}

func (d *LinkedList) IntoArray() (arr []int) {
	for iter := d.Head; iter != nil; iter = iter.next {
		arr = append(arr, iter.key)
	}
	return
}

func (d *LinkedList) Revert() {
	curr := d.Head
	var prev *llNode

	for curr != nil {
		next := curr.next

		curr.next = prev
		prev = curr

		curr = next
	}

	d.Head = prev
}
