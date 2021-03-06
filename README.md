### RandPool

RandPool is a benchmark to test if we can get faster than rand instance if we use a pool. The result? No. It also is not thread safe [*] :(

However! If we are looking to have unique values throughout a distributed system, we can make use of a type of Lamport Clock. It is fast and guaranteed unique.


```
$ go test -bench=.
BenchmarkRandPool-8            	50000000	        29.1  ns/op
BenchmarkRandInstance-8        	100000000	        15.4  ns/op
BenchmarkRandPackage-8         	50000000	        28.4  ns/op
BenchmarkFastRand-8            	10000000	       213    ns/op
BenchmarkClock-8               	200000000	         8.05 ns/op
PASS
ok  	github.com/sethgrid/randpool	9.335s
```

[*] note that RandInstance and RandPool below are not actually thread safe. You can verify this by running the benchmark with `-race` and unskipping the race test in the benchmark. This is because `rand.NewSource` is not thread safe.
