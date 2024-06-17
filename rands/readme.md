## Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/ulib/rands
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_Rands_Std_V2-20    	  591513	      1934 ns/op	    2128 B/op	       2 allocs/op
Benchmark_Rands_Std_V1-20    	  594921	      2021 ns/op	    2128 B/op	       2 allocs/op
Benchmark_Rands-20           	  661988	      1713 ns/op	    2128 B/op	       2 allocs/op
Benchmark_CRands-20          	  533282	      2360 ns/op	    2129 B/op	       2 allocs/op
PASS
ok  	github.com/3JoB/ulib/rands	5.058s
```