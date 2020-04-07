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

	if elem == d.head && elem == d.tail {

		d.head = nil
		d.tail = nil
		d.size--
		return value

	} else if elem == d.head {
		val := elem.value

		front := elem.next
		front.prev = nil
		elem = nil
		d.head = front
		d.size--
		return val

	} else if elem == d.tail {
		val := elem.value

		back := elem.prev
		back.next = nil
		d.tail = back
		d.size--
		return val

	} else {
		val := elem.value
		back := elem.prev
		front := elem.next
		back.next = elem.next
		front.prev = elem.prev
		elem = nil
		d.size--
		return val
	}

}

// PeekFront returns the Head node of the list.
func (d *DoubleList) PeekFront() interface{} {
	if d.size == 0 {
		return nil
	}
	return d.head.value
}

// PeekBack returns the Tail node of the list.
func (d *DoubleList) PeekBack() interface{} {
	if d.size == 0 {
		return nil
	}
	return d.tail.value
}

// Len returns the length of the list.
func (d *DoubleList) Len() int {
	return d.size
}

/* To be used only in development environment.
func (d *DoubleList) Iterate() {

	fmt.Println("Size: ", d.size)
	for l := d.head; l != nil; l = l.next {
		fmt.Println(l)
	}
}
*/
