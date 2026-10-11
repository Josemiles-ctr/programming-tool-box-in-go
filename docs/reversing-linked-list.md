# `datastructures/reversing_linked_list.go`

[`datastructures/reversing_linked_list.go`](../datastructures/reversing_linked_list.go) implements `ReverseLinkedList` for the package's singly linked `Node` type. It iterates through the input, saves each node's original successor, points the node back to the already-reversed prefix, and advances. It returns the new head; an empty input returns `nil`.

The function changes the links in place, uses O(1) additional space, and takes O(n) time for a list of n nodes. The caller should use the returned pointer as the new head.
