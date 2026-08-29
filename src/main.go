package main
import (
	"fmt"
   "goProgramming/datastructures"

)
func main() {
	head := datastructures.LinkedList()
	current := head
	fmt.Println("Original Linked List:")
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	reversedHead := datastructures.ReverseLinkedList(head)
	fmt.Println("\nReversed Linked List:")
	current = reversedHead
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
}
