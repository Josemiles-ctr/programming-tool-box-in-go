package datastructures

type Node struct {
	Data int
	Next  *Node
}

func LinkedList() *Node {
	head := &Node{Data: 1}
	second := &Node{Data: 2}
	third := &Node{Data: 3}
	fourth := &Node{Data: 4}

	head.Next = second
	second.Next = third
	third.Next = fourth
	
	return head
}
