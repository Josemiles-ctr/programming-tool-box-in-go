# `main.go`

[`main.go`](../main.go) is the executable entry point for the repository. Its `main` function creates an `algorithms.Cache` with capacity four, initializes it with the sample domains, and prints the list. It then adds `ticktok.com` and `josemiles.com`, printing after each insertion, and adds the same `tiktok` node again to demonstrate duplicate handling.

The cache prints most-recent entries first. Since it starts full, each new distinct value evicts the least-recent entry. The duplicate call moves the existing value to the front when needed without increasing the count. The commented `goBasics.Person` lines are an inactive example. See [LRU cache](LRU.md) for the cache's implementation details and limitations.
