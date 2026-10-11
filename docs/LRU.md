# LRU Cache

The LRU cache in [`algorithms/lru.go`](../algorithms/lru.go) stores string values in a doubly linked list. The list is ordered from most recently used to least recently used: `Head` is the most recent entry, and the final node is the eviction candidate.

## Current implementation

`Cache` exposes its `Capacity`, `CurrentCount`, and `Head`. Each `Node` holds a string in `Data` and links to its neighboring nodes through `Prev` and `Next`.

`AddCacheEntry` is the cache's only access/update operation:

- If a node with the same `Data` is already present, it is moved to the head, unless it is already there. The existing node is retained; the passed node is not inserted.
- If the value is new and the cache is below capacity, the count is incremented and the new node is added at the head.
- If the cache is at capacity, the tail is detached and the new node is added at the head. The count stays the same.

`AlreadyExists` walks the list to find both a matching node and the tail. Consequently, adding an entry takes O(n) time and uses O(1) extra space. The list itself uses O(n) space for the entries. `PrintLinkedList` prints values from most to least recent, separated by arrows.

## Initialization and limitations

`InitializeCache` inserts four hard-coded entries—`google.com`, `youtube.com`, `amazon.com`, and `facebook.com`—in that order of calls. Since every insertion goes to the head, the resulting list is `facebook.com -> amazon.com -> youtube.com -> google.com`. Set `Capacity` before calling it; the method does not set a capacity itself.

The current API has no lookup method, so it does not track reads as use; recency changes only when `AddCacheEntry` is called. Capacity and count are public and are not validated. In particular, eviction assumes a non-empty list with a non-nil tail predecessor, so zero or negative capacity and inconsistent count/list state are not handled safely. This implementation is a learning example rather than a production-ready cache.
