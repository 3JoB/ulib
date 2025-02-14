Benchmark

## New
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/ulib/fsutil/writer
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_Ulib_Write-20                                 	    1772	    650801 ns/op	    4717 B/op	       4 allocs/op
Benchmark_Ulib_Lite_Write-20                            	    3265	    386735 ns/op	    4712 B/op	       4 allocs/op
Benchmark_UlibWriter_Write-20                           	  140640	      7806 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Lite_Write-20                      	  134116	      9836 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Strings_Write-20                   	  144152	      9308 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_Strings_Lite_Write-20              	259474774	         4.815 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Write-20                         	  168164	      9195 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Lite_Write-20                    	  132193	      9450 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Strings_Write-20                 	  163704	      9435 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibOSWriter_Strings_Lite_Write-20            	  366168	      2889 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_MaxBuffer_Strings_Write-20         	  140949	      9489 ns/op	       0 B/op	       0 allocs/op
Benchmark_UlibWriter_MaxBuffer_Strings_Lite_Write-20    	263471228	         4.541 ns/op	       0 B/op	       0 allocs/op
Benchmark_Basic_Write-20                                	 1226612	       961.0 ns/op	    6528 B/op	       1 allocs/op
Benchmark_Basic_Lite_Write-20                           	189490268	         6.090 ns/op	       1 B/op	       1 allocs/op
PASS
coverage: 63.0% of statements
ok  	github.com/3JoB/ulib/fsutil/writer	23.739s
```

## Old
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