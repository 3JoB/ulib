Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/ulib/fsutil/writer
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
Benchmark_Ulib_Write-32                                 	    2554	    461757 ns/op	    4808 B/op	       4 allocs/op
Benchmark_Ulib_Lite_Write-32                            	    2606	    472146 ns/op	    4808 B/op	       4 allocs/op
Benchmark_UlibWriter_Write-32                           	   89488	     12976 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Lite_Write-32                      	  123360	     10497 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Strings_Write-32                   	  125884	     10036 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Strings_Lite_Write-32              	95578688	        11.08 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Write-32                         	   96375	     10968 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Lite_Write-32                    	  126484	      9635 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Strings_Write-32                 	  128119	      9648 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Strings_Lite_Write-32            	  395128	      2533 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_MaxBuffer_Strings_Write-32         	   82396	     12732 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_MaxBuffer_Strings_Lite_Write-32    	106982149	        11.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_Basic_Write-32                                	  575473	      1933 ns/op	    6528 B/op	       1 allocs/op
Benchmark_Basic_Lite_Write-32                           	59480732	        20.30 ns/op	       1 B/op	       1 allocs/op
PASS
ok  	github.com/3JoB/ulib/fsutil/writer	23.202s
```