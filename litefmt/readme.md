# Benchmark
## New
```
goos: windows
goarch: amd64
pkg: github.com/3JoB/ulib/litefmt
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_String_For-20          	  201957	      5602 ns/op	   24328 B/op	      95 allocs/op
Benchmark_Strings_Join-20        	 2158663	       571.0 ns/op	     576 B/op	       1 allocs/op
Benchmark_LiteFMT_Sprint-20      	 3240400	       375.4 ns/op	     480 B/op	       1 allocs/op
Benchmark_LiteFMT_SprintP-20     	 4224103	       291.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_FMT_Sprint-20          	  266916	      4522 ns/op	    2201 B/op	      98 allocs/op
Benchmark_LiteFMT_Sprintln-20    	 3224475	       388.1 ns/op	     512 B/op	       1 allocs/op
Benchmark_FMT_Sprintln-20        	  260229	      4520 ns/op	    2201 B/op	      98 allocs/op
Benchmark_L_LITEFMT_Sprint-20    	32310177	        35.29 ns/op	      48 B/op	       1 allocs/op
Benchmark_L_FMT_Sprint-20        	 6072588	       179.9 ns/op	     104 B/op	       4 allocs/op
PASS
coverage: 61.1% of statements
ok  	github.com/3JoB/ulib/litefmt	12.932s
```

## Old
```
goarch: amd64
pkg: github.com/3JoB/ulib/litefmt
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
Benchmark_LiteFMT_Sprint-20         	 2609953	       484.9 ns/op	    1016 B/op	       7 allocs/op
Benchmark_LiteFMT_PSprint_New-20       	 3254678	       383.7 ns/op	     480 B/op	       1 allocs/op
Benchmark_LiteFMT_PSprintP_New-20      	 4159430	       303.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_LiteFMT_VSprint-20        	 1746712	       678.8 ns/op	    1920 B/op	       5 allocs/op
Benchmark_LiteFMT_LSprint-20        	 2286727	       517.3 ns/op	    1016 B/op	       7 allocs/op
Benchmark_LiteFMT_PSprint_OLD-20        	 2296669	       581.3 ns/op	    1016 B/op	       7 allocs/op
Benchmark_LITEFMT_TSprint-20        	 2295609	       526.7 ns/op	    1496 B/op	       8 allocs/op
Benchmark_FMT_Sprint-20             	  223491	      5252 ns/op	    2201 B/op	      98 allocs/op
Benchmark_LiteFMT_Sprintln-20       	 2610094	       450.3 ns/op	    1016 B/op	       7 allocs/op
Benchmark_FMT_Sprintln-20           	  234364	      5145 ns/op	    2201 B/op	      98 allocs/op
Benchmark_L_LITEFMT_Sprint-20       	31340511	        38.41 ns/op	      64 B/op	       2 allocs/op
Benchmark_L_LiteFMT_PSprint_New-20     	34987258	        35.54 ns/op	      48 B/op	       1 allocs/op
Benchmark_L_LiteFMT_PSprintP_New-20    	63030175	        17.90 ns/op	       0 B/op	       0 allocs/op
Benchmark_L_FMT_Sprint-20           	 6037972	       197.7 ns/op	     104 B/op	       4 allocs/op
PASS
ok  	github.com/3JoB/ulib/litefmt	22.262s
```