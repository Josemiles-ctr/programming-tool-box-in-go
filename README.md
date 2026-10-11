# goProgramming

A small collection of Go examples for learning basic structs, linked lists, and algorithms. The examples are organized by package so you can explore each implementation independently.

## Project files

Every Go source file has a reference document in the [`docs/`](docs/) folder. Read the matching page for an explanation of the file's purpose, current behavior, and any implementation limitations.

| Source file | Reference documentation |
| --- | --- |
| [`main.go`](main.go) | [Main program](docs/main.md) |
| [`goBasics/structs.go`](goBasics/structs.go) | [Structs basics](docs/structs.md) |
| [`datastructures/linked_list.go`](datastructures/linked_list.go) | [Singly linked list](docs/linked-list.md) |
| [`datastructures/reversing_linked_list.go`](datastructures/reversing_linked_list.go) | [Reversing a linked list](docs/reversing-linked-list.md) |
| [`datastructures/product_except.go`](datastructures/product_except.go) | [Product except one position](docs/product-except.md) |
| [`algorithms/lru.go`](algorithms/lru.go) | [LRU cache](docs/LRU.md) |

## Run

Run the example program from the repository root:

```sh
go run .
```

The module name is `goProgramming`; the source packages can also be imported by their module paths, such as `goProgramming/algorithms`.

## Learning focus

The repository is intended for practicing Go fundamentals and examining straightforward implementations of common data structures and algorithms. Some functions are incomplete or have narrow input assumptions; their reference pages describe what the current code does.
