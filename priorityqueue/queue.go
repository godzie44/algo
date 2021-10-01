package priorityqueue

import (
	"errors"
	"math"
)

type PriorityQueue struct {
	a         []*QueueElement
	greaterFn func(key1, key2 int) bool
	typ       QueueType
}

type QueueType int

const (
	_ QueueType = iota
	MinQueue
	MaxQueue
)

func NewPriorityQueue(queueType QueueType) *PriorityQueue {
	var greaterFn func(key1, key2 int) bool
	switch queueType {
	case MaxQueue:
		greaterFn = func(key1, key2 int) bool {
			return key1 > key2
		}
	case MinQueue:
		greaterFn = func(key1, key2 int) bool {
			return key1 < key2
		}
	default:
		panic("unknown type")
	}

	return &PriorityQueue{
		greaterFn: greaterFn,
		typ:       queueType,
	}
}

type QueueElement struct {
	Key   int
	Value interface{}
}

func (p *PriorityQueue) MaxOrMin() *QueueElement {
	return p.a[0]
}

var ErrEmptyQueue = errors.New("empty queue")

func (p *PriorityQueue) ExtractMaxOrMin() (*QueueElement, error) {
	if len(p.a) < 1 {
		return nil, ErrEmptyQueue
	}

	max := p.a[0]
	p.a[0] = p.a[len(p.a)-1]
	p.a = p.a[:len(p.a)-1]

	p.heapify(p.a, 0)
	return max, nil
}

func (p *PriorityQueue) IncreaseKey(i, key int) error {
	if !p.greaterFn(key, p.a[i].Key) {
		return errors.New("new key less then current")
	}

	parent := func(i int) int {
		return (i - 1) / 2
	}

	p.a[i].Key = key
	for i > 0 && !p.greaterFn(p.a[parent(i)].Key, p.a[i].Key) {
		p.a[i], p.a[parent(i)] = p.a[parent(i)], p.a[i]
		i = parent(i)
	}

	return nil
}

func (p *PriorityQueue) Insert(el *QueueElement) {
	k := el.Key
	if p.typ == MaxQueue {
		el.Key = math.MinInt64
	} else {
		el.Key = math.MaxInt64
	}
	p.a = append(p.a, el)
	_ = p.IncreaseKey(len(p.a)-1, k)
}

func (p *PriorityQueue) heapify(arr []*QueueElement, ind int) {
	for {
		left := ind*2 + 1
		right := ind*2 + 2

		largest := ind
		if left <= len(arr)-1 && p.greaterFn(arr[left].Key, arr[ind].Key) {
			largest = left
		}
		if right <= len(arr)-1 && p.greaterFn(arr[right].Key, arr[largest].Key) {
			largest = right
		}

		if largest == ind {
			return
		}

		arr[ind], arr[largest] = arr[largest], arr[ind]
		ind = largest
	}
}

func (p *PriorityQueue) ForEach(fn func(ind int, el *QueueElement)) {
	for i, el := range p.a {
		fn(i, el)
	}
}
