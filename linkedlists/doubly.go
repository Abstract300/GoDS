// Inspired by: https://github.com/emirpasic/gods

package linkedlists

// Node is a node in DLL
type Node struct {
	// next and prev are pointers to next, prev nodes.
	next, prev *Node
	// value is the value of the node
	value interface{}
}

// DoubleList is a list that's made of Node(s)
type DoubleList struct {
	// head and tail represent the begnning and ending node of the list.
	head, tail *Node
	// size represents the total number of nodes in the list.
	size int
}

// MakeList returns an empty DoubleList
func MakeList() *DoubleList {
	return &DoubleList{}
}

// AddBack adds elements at the end(Tail) of the list.
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

// AddFront adds elements at the beginning(Head) of the list.
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

// Remove removes a Node from the list at the given index.
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

// PeekFront returns the Head node of the list.
func (d *DoubleList) PeekFront() *Node {
	return d.head
}

// PeekBack returns the Tail node of the list.
func (d *DoubleList) PeekBack() *Node {
	return d.tail
}

// Len returns the length of the list.
func (d *DoubleList) Len() int {
	return d.size
}
