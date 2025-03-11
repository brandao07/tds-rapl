# Results

`measurements_<x>_<y>.csv`

- x: with or without flags (`wf or wof`)
- y: number of iterations

- With Flags: `go build -o fib fib.go`
- Without Flags: `go build -gcflags="all=-N -l" -o fib fib.go`

# References

- [Fibonacci Sequence with Go](https://opiticalvin.medium.com/fibonacci-sequence-with-go-3eecb9ff2f57)
- [Efficient Fibonacci Series Algorithm in Go](https://medium.com/@nagarjun_nagesh/efficient-fibonacci-series-algorithm-in-go-ababd4ff985e)
- [Exploring Sorting Algorithms: In Golang](https://edwinsiby.medium.com/common-sorting-methods-in-data-structure-with-golang-d02ec6b965c8)