package linkedlists

type SingleNode struct {
	Value interface{}
	Next  *SingleNode
}

func InitSLL(data interface{}) SingleNode {
	return SingleNode{
		Value: data,
		Next:  nil,
	}
}

func (sll *SingleNode) SingleLLTail() SingleNode {
	var retNode SingleNode
	for s := sll; s != nil; s = s.Next {
		if s.Next == nil {
			retNode = *s
		}
	}
	return retNode
}

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

func (sll *SingleNode) NextNode() SingleNode {
	if sll.Next != nil {
		return *sll.Next
	}
	return SingleNode{}
}
