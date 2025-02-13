## ArrayPool
![go build workflow](https://github.com/j0nimost/arraypool/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/j0nimost/arraypool)](https://goreportcard.com/report/github.com/j0nimost/arraypool)

A fast implementation of generic reusable buffers. Uses `sync.pool` under the hood.

When creating and disposing arrays is an expensive operation. `ArrayPool[T any]` allows you to
create buffers that can be used across the application lifecycle. 

- Installation
  `go get github.com/j0nimost/arraypool/v2`
- Supported Version 
  go >= 1.21
- Benchmark
```txt
goos: linux
goarch: amd64
pkg: arraypool
cpu: AMD Ryzen 7 5800U with Radeon Graphics


BenchmarkArrayPool-4    12008463                95.50 ns/op
BenchmarkSlice-4         2217805               545.4 ns/op

```
The above benchmark is included in the test, it compares a normal slice vs an arraypool

### Usage
```go
import ap "github.com/j0nimost/arraypool"
// initialize an array pool instance, make it global 
var apool *ap.ArrayPool[int] = ap.NewPool[int]() 

func methodThatCreateAlotofSlices() {
  // get an existing or new buffer from a global declaration
  apbuffer := apool.Get()
  // ... do work with the array
  //
  // write to buffer 
  apBuffer.Write(5)
  // put back in the pool after use, also a flag to resize
  apool.Put(apbuffer, true)// the flag `true` means the underlying buffer's length is set to 0
}
```

The `ArrayPoolBuffer[T any]` comes with some helpers;
#### API
```go
Write(val T) // writes val to the buffer
ClearAll() // clears the values of the buffer
Resize() // resizes the length of the buffer to zero
  
```
### Author
John Nyingi
