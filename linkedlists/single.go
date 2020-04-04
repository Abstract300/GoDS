package linkedlists

type SingleNode struct {
	value interface{}
	next  *SingleNode
}

func MakeSingleLL(data interface{}) SingleNode {
	return SingleNode{
		value: data,
		next:  nil,
	}
}

func (sll *SingleNode) SingleLLTail() SingleNode {
	for {
		if sll.next == nil {
			return *sll
		}
		sll = sll.next
	}
}

func (sll *SingleNode) SingleLLAppend(data interface{}) {
	node := sll.SingleLLTail()

	node.next = &SingleNode{
		value: data,
		next:  nil,
	}

	sll.next = &node
}
