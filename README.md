### RandPool

RandPool is a benchmark to test if we can get faster than rand instance if we use a pool. The result? No.

```
$ go test -bench=.
BenchmarkRandPool-8       	20000000	        65.3 ns/op
BenchmarkRandInstance-8   	100000000	        16.0 ns/op
BenchmarkRandPackage-8    	50000000	        29.9 ns/op
BenchmarkFastRand-8       	10000000	       214 ns/op
PASS
ok  	github.com/sethgrid/randpool	7.005s
```