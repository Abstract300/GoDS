// Inspired by: https://github.com/emirpasic/gods

package linkedlists

type Node struct {
	next, prev *Node
	value      interface{}
}

type DoubleList struct {
	head, tail *Node
	size       int
}

func MakeList() *DoubleList {
	return &DoubleList{}
}

func (d *DoubleList) AddBack(data interface{}) {
	newElement := &Node{value: data, prev: d.tail}

	if d.size == 0 {
		d.head = newElement
		d.tail = newElement
	} else {
		d.tail.next = newElement
		d.tail = newElement
	}

	d.size++
}

func (d *DoubleList) AddFront(data interface{}) {
	newElement := &Node{value: data, next: d.head}

	if d.size == 0 {
		d.head = newElement
		d.tail = newElement
	} else {
		d.head.prev = newElement
		d.head = newElement
	}
	d.size++
}

func (d *DoubleList) Remove(index int) interface{} {
	var value interface{}
	var counter int
	var elem *Node

	if index < 0 || index >= d.size {
		return value
	}

	for l := d.head; l != nil; l = l.next {
		if counter == index {
			elem = l
		}
		counter++
	}

	if elem == d.head {
		d.head = d.head.next
		d.head.prev = nil
		return elem.value

	} else if elem == d.tail {
		d.tail = d.tail.prev
		d.tail.next = nil
		return elem.value

	} else {
		elem.prev.next = elem.next
		elem.next = elem.prev
		return elem.value
	}

}
func (d *DoubleList) PeekFront() *Node {
	return d.head
}

func (d *DoubleList) PeekBack() *Node {
	return d.tail
}

func (d *DoubleList) Len() int {
	return d.size
}
