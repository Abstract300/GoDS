package main

import (
	"fmt"

	"github.com/abstract300/ds/linkedlists"
)

func main() {
	list := linkedlists.MakeSingleLL(20)
	list.SingleLLAppend(40)
	list.SingleLLAppend(50)
	list.SingleLLAppend(60)
	list.SingleLLAppend(70)
	fmt.Println(list)
	fmt.Println(list.SingleLLTail())
}
