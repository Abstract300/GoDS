package main

import (
	"fmt"

	"github.com/abstract300/ds/linkedlists"
)

func main() {
	list := linkedlists.InitSLL(20)
	list.SingleLLAppend(40)
	list.SingleLLAppend(50)
	list.SingleLLAppend(60)
	list.SingleLLAppend(70)
	list.SingleLLAppend(80)
	list.SingleLLAppend(90)

	for l := &list; l != nil; l = l.Next {
		fmt.Println(l)
	}
	fmt.Println()

	list.SingleLLDelete(0)
	for l := &list; l != nil; l = l.Next {
		fmt.Printf("Next to %v is %v\n", l, l.NextNode())
	}

}
