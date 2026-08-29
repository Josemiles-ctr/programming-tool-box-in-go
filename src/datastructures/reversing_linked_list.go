package datastructures

func ReverseLinkedList(node *Node) *Node {
	current :=node
	var next *Node
	var previous *Node

  for current!=nil{
	 next=current.Next	
	 current.Next=previous;

		previous=current;
		current=next
	}
	return previous
}
