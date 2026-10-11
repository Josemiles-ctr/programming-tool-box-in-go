# `datastructures/product_except.go`

[`datastructures/product_except.go`](../datastructures/product_except.go) defines `ProductExcept`, which takes a fixed-size array of four integers and an index `out`. It multiplies every array element whose index differs from `out`, then returns that product.

The function runs in O(4) time (constant for this fixed array size) and uses O(1) additional space. It does not validate `out`: an index outside `0` through `3` excludes no element, so the result is the product of all four values. A zero among the included values makes the result zero.
