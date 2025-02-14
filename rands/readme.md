## Benchmark

### New
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/ulib/rands
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkRandBase/stdV1_large-20         	441347654	         2.783 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandBase/stdV2_large-20         	444422386	         2.731 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandBase/frand_large-20         	455896219	         2.749 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandBase/pg_large-20            	433610132	         2.732 ns/op	       0 B/op	       0 allocs/op
BenchmarkRand/stdV1_large-20             	  135700	      8652 ns/op	   12288 B/op	       2 allocs/op
BenchmarkRand/stdV2_large-20             	  138212	      8506 ns/op	   12288 B/op	       2 allocs/op
BenchmarkRand/frand_large-20             	  135483	      9139 ns/op	   12297 B/op	       2 allocs/op
BenchmarkRand/pg_large-20                	  165686	      7487 ns/op	   12288 B/op	       2 allocs/op
PASS
coverage: 97.1% of statements
ok  	github.com/3JoB/ulib/rands	11.293s
```

### Old
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