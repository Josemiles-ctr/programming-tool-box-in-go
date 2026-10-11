package algorithms

import (
	"fmt"
)
type Cache struct {
	Capacity int
	CurrentCount int
	Head *Node
}

type Node struct {
	Data string
	Prev *Node
	Next  *Node
}

func (c *Cache)InitializeCache() *Cache {
	head := &Node{Data: "google.com"}
	second := &Node{Data: "youtube.com"}
	third := &Node{Data: "amazon.com"}
	fourth := &Node{Data: "facebook.com"}

	c.Head = head
	c.CurrentCount++
	c.AddCacheEntry(second)
	c.AddCacheEntry(third)
	c.AddCacheEntry(fourth)
	
	return c
}
func (c *Cache) PrintLinkedList(){
 for node:=c.Head; node!=nil; node=node.Next {
	fmt.Printf("%s ->", node.Data)
	}
	println()
}

func (c *Cache) AddCacheEntry(newNode *Node) *Cache {
	existing, last := c.AlreadyExists(newNode)
	if existing != nil {
		if existing == c.Head {
			return c
		}

		existing.Prev.Next = existing.Next
		if existing.Next != nil {
			existing.Next.Prev = existing.Prev
		}
		existing.Prev = nil
		existing.Next = c.Head
		c.Head.Prev = existing
		c.Head = existing
		return c
	}

	if c.CurrentCount == c.Capacity {
		last.Prev.Next = nil
		last.Prev = nil
	} else {
		c.CurrentCount++
	}

	newNode.Prev = nil
	newNode.Next = c.Head
	if c.Head != nil {
		c.Head.Prev = newNode
	}
	c.Head = newNode
	return c
}

func (c *Cache) AlreadyExists(newNode *Node) (*Node, *Node){
	repeated, last :=(*Node)(nil), (*Node)(nil)
	for node:=c.Head; node!=nil; node=node.Next {
		if node.Data == newNode.Data {
			repeated=node		
		}
		last=node
	}
	return repeated, last
}