# `datastructures/linked_list.go`

[`datastructures/linked_list.go`](../datastructures/linked_list.go) defines a singly linked `Node` containing a string and a `Next` pointer. `LinkedList` constructs and returns a four-node list containing `google.com`, `youtube.com`, `amazon.com`, and `facebook.com` in that order. `PrintLinkedList` walks a list from its supplied head and prints each value followed by an arrow.

`AddCache` currently returns the original head without changing the list. It is a placeholder and does not add the supplied domain. This package's singly linked list is separate from the doubly linked LRU implementation documented in [LRU cache](LRU.md).
