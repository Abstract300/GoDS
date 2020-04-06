package linkedlists

// SingleNode represents a Node in SLL.
type SingleNode struct {
	Value interface{}
	Next  *SingleNode
}

// InitSLL initialises and returns an SLL.
func InitSLL(data interface{}) SingleNode {
	return SingleNode{
		Value: data,
		Next:  nil,
	}
}

// SingleLLTail returns the tail of the SLL.
func (sll *SingleNode) SingleLLTail() SingleNode {
	var retNode SingleNode
	for s := sll; s != nil; s = s.Next {
		if s.Next == nil {
			retNode = *s
		}
	}
	return retNode
}

// SingleLLAppend adds a node to the end of the list.
func (sll *SingleNode) SingleLLAppend(data interface{}) {
	for s := sll; s != nil; s = s.Next {
		if s.Next == nil {
			s.Next = &SingleNode{
				Value: data,
				Next:  nil,
			}
			break
		}
	}

}

// SingleLLDelete deletes a node at the given index.
func (sll *SingleNode) SingleLLDelete(position int) {

	var counter int

	if position == 0 {
		sll.Value = sll.Next.Value
		sll.Next = sll.Next.Next
	}

	for s := sll; s != nil; s = s.Next {
		if counter == position-1 {
			if s.Next != nil {
				s.Next = s.Next.Next
				break
			}
		}
		counter++
	}

}

// NextNode returns the consecutive node.
func (sll *SingleNode) NextNode() SingleNode {
	if sll.Next != nil {
		return *sll.Next
	}
	return SingleNode{}
}
