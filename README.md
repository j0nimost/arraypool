## ArrayPool
![go build workflow](https://github.com/j0nimost/arraypool/actions/workflows/go.yml/badge.svg)

A fast implementation of generic reusable buffers. Uses `sync.pool` under the hood.

When creating and disposing arrays is an expensive operation. `ArrayPool[T any]` allows you to
create buffers that can be used across the application lifecycle. 

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
// initialize an array pool instance, make it global 
var ap ArrayPool[int] = ArrayPool[int]{}

func methodThatCreateAlotofSlices() {

  // create an instance of a buffer do this once;
  apbuffer:= ap.New()
  // or 
  // get an exisiting buffer
  apbuffer := ap.Get()
  // put back in the pool after use, also a flag to resize
  // ... do work with the array
  //
  // 
  ap.Put(apbuffer, true)// the flag `true` means the underlying buffer's length is set to 0
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
