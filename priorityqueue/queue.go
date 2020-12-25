package priorityqueue

import (
	"errors"
	"math"
)

type PriorityQueue struct {
	a []*QueueElement
}

type QueueElement struct {
	Key   int
	Value interface{}
}

func (p *PriorityQueue) Maximum() *QueueElement {
	return p.a[0]
}

func (p *PriorityQueue) ExtractMaximum() (*QueueElement, error) {
	if len(p.a) < 1 {
		return nil, errors.New("empty queue")
	}

	max := p.a[0]
	p.a[0] = p.a[len(p.a)-1]
	p.a = p.a[:len(p.a)-1]

	maxHeapify(p.a, 0)
	return max, nil
}

func (p *PriorityQueue) increaseKey(i, key int) error {
	if key < p.a[i].Key {
		return errors.New("new key less then current")
	}

	parent := func(i int) int {
		return (i - 1) / 2
	}

	p.a[i].Key = key
	for i > 0 && p.a[parent(i)].Key < p.a[i].Key {
		p.a[i], p.a[parent(i)] = p.a[parent(i)], p.a[i]
		i = parent(i)
	}

	return nil
}

func (p *PriorityQueue) Insert(el *QueueElement) {
	k := el.Key
	el.Key = math.MinInt64
	p.a = append(p.a, el)
	_ = p.increaseKey(len(p.a)-1, k)
}

func maxHeapify(arr []*QueueElement, ind int) {
	for {
		left := ind*2 + 1
		right := ind*2 + 2

		largest := ind
		if left <= len(arr)-1 && arr[left].Key > arr[ind].Key {
			largest = left
		}
		if right <= len(arr)-1 && arr[right].Key > arr[largest].Key {
			largest = right
		}

		if largest == ind {
			return
		}

		arr[ind], arr[largest] = arr[largest], arr[ind]
		ind = largest
	}
}
