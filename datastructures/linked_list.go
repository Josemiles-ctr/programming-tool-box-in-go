package datastructures

import (
	"fmt"
)

type Node struct {
	Data string
	Next  *Node
}

func LinkedList() *Node {
	head := &Node{Data: "google.com"}
	second := &Node{Data: "youtube.com"}
	third := &Node{Data: "amazon.com"}
	fourth := &Node{Data: "facebook.com"}

	head.Next = second
	second.Next = third
	third.Next = fourth
	
	return head
}
func PrintLinkedList(head *Node){
 for head!=nil{
	fmt.Printf("%s ->", head.Data)
	head=head.Next
	}
}
func AddCache(head *Node, domain string) *Node{

	return head;
}
