package main

import (
	"fmt"

	"test.com/DSandAlgo/DataStructure/linkedlist"
)

func main() {
	myList := linkedlist.LinkedList{}
	node1 := &linkedlist.Node{Data: 45, Next: nil}
	node2 := &linkedlist.Node{Data: 46, Next: nil}
	node3 := &linkedlist.Node{Data: 47, Next: nil}
	myList.Append(node1)
	myList.Append(node2)
	myList.Append(node3)
	fmt.Println(myList)
}
