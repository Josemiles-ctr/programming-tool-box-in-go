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
}

func (c *Cache) AddCacheEntry(newNode *Node) *Cache {
	
	newNode.Prev=nil
	referenceHead:=c.Head
	if c.CurrentCount == c.Capacity {
		for c.Head.Next !=nil{
			c.Head=c.Head.Next
		}
		c.Head.Prev.Next=nil;
		newNode.Next=referenceHead
		referenceHead.Prev=newNode
		c.Head=newNode
		return c
	}else{
		newNode.Next=c.Head
		c.CurrentCount++
		c.Head.Prev=newNode
		c.Head=newNode
		return c
	}
}
