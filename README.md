# Unified Rendezvous Hash (hrw)

This is implementation of [rendezvous hashing](https://en.wikipedia.org/wiki/Rendezvous_hashing)
inspired by [github.com/nspcc-dev/hrw](https://github.com/nspcc-dev/hrw).

## Description
Rendezvous hashing allows to choose object from set of objects in synchronous way. This can be useful
for distributed systems: nodes with same pivot value can independently take same decision from the same 
pool of decisions. This library provides unified interface for work with HRW. You choose hash function
and you can shuffle slice of any objects, which implement `Raw() []byte` function. 

In contrast to others implementation, there used simple approach for hashing routine. One of the main 
issues of HRW implementations is a uniform distribution of results. If weights are assigned by 
hash of objects then there is a possibility that some hashes will be closer to each other than others.

Instead of hash this library use position in sorted array of objects as a weight. Hash from pivot object
is unsigned number. If hash function is good enough, this number will be distributed with uniform distribution 
between `0` and `MaxInt`. We can use a reminder from division to number of object as an result of hashing routine.
These simplifications allow to see influence of hash function. This way library works more transparent, 
which is good for audit and debugging in general. 

## Install

```
$ go get github.com/AlexVanin/unihrw
```

## Usage

You can choose any hash function, which implement `hash.Hash32` or `hash.Hash64` interface. It is recommended
to use murmur3 hash library. As an input `HrwSort` functions takes `[][]byte`, `[]string`, `[]int` or slice 
of object with implemented `Rawer` interface.

## Example

```go
package main

import (
	"fmt"
	"github.com/AlexVanin/unihrw"
	"github.com/spaolacci/murmur3"
)

func main() {
	nodes := []string{
	    "12.43.5.123:2090",
		"10.15.1.84:8888",
		"192.168.5.44:4000",
		"10.17.0.3:2019",
		"172.0.1.2:44412",
    }

	hash := murmur3.New32()
	key1 := []byte("aabbcc")
	err := unihrw.HrwSort32(nodes, key1, hash)
	if err == nil {
		fmt.Println(nodes[0])
	}
}
```




## Benchmarks

```
BenchmarkMurMur32Objects/10-8         	  500000	      3019 ns/op	     819 B/op	      18 allocs/op
BenchmarkMurMur32Strings/10-8         	 3000000	       559 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur32Strings/100-8        	  200000	      6829 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur32Strings/1000-8       	  200000	      6862 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur32Bytes/10-8           	 1000000	      1412 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur32Bytes/100-8          	  100000	     17122 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur32Bytes/1000-8         	  100000	     16699 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur32Ints/10-8            	 3000000	       443 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur32Ints/100-8           	  500000	      3537 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur32Ints/1000-8          	  500000	      3543 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Objects/10-8         	  500000	      2972 ns/op	     820 B/op	      18 allocs/op
BenchmarkMurMur64Objects/100-8        	   20000	     62599 ns/op	    5726 B/op	     113 allocs/op
BenchmarkMurMur64Objects/1000-8       	   20000	     62503 ns/op	    5728 B/op	     113 allocs/op
BenchmarkMurMur64Strings/10-8         	 2000000	       614 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Strings/100-8        	  200000	      7817 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Strings/1000-8       	  200000	      7878 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Bytes/10-8           	 1000000	      1381 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur64Bytes/100-8          	  100000	     18006 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur64Bytes/1000-8         	  100000	     17453 ns/op	     224 B/op	       6 allocs/op
BenchmarkMurMur64Ints/10-8            	 3000000	       488 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Ints/100-8           	  300000	      4248 ns/op	     128 B/op	       4 allocs/op
BenchmarkMurMur64Ints/1000-8          	  300000	      4147 ns/op	     128 B/op	       4 allocs/op
```

## ToDo

- Speed of HRW is depended on speed of sorting routine. Objects with `Rawer` interface sorting slowly,
so maybe there is a room for optimizations.
- Tests are not covered all the code

