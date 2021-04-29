# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test resule for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=10s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJsonStruct-8                    	 1349085	      8670 ns/op	    1144 B/op	      25 allocs/op
Benchmark__Marshal__GoStdJsonStruct-8                    	 1776642	      6704 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf_blob-8               	   10000	   1059023 ns/op	  414676 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf-8                    	  977202	     12420 ns/op	    4512 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJsonMapItf-8                    	  739980	     16309 ns/op	    5865 B/op	     121 allocs/op
Benchmark_Unmarshal_JsoniterStruct-8                     	 1823509	      6484 ns/op	    1720 B/op	      56 allocs/op
Benchmark__Marshal__JsoniterStruct-8                     	 1776775	      6712 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_EasyjsonStruct-8                     	 3170910	      3765 ns/op	     784 B/op	      19 allocs/op
Benchmark__Marshal__EasyjsonStruct-8                     	 5010370	      2395 ns/op	    1240 B/op	       5 allocs/op
Benchmark_Unmarshal_JsoniterMapItf_blob-8                	   12758	    922637 ns/op	  457039 B/op	   13302 allocs/op
Benchmark_Unmarshal_JsoniterMapItf-8                     	 1310109	      9094 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__JsoniterMapItf-8                     	  738163	     16383 ns/op	    5865 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                           	 5125300	      2343 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8         	 1934654	      6133 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8              	 1403001	      8508 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8              	  971823	     12371 ns/op	    3665 B/op	     197 allocs/op
Benchmark____Get____Jsoniter_blob_ReadOneChain-8         	   29994	    403199 ns/op	  104264 B/op	    6788 allocs/op
Benchmark_Unmarshal_Jsonparser_Full_Blob-8               	   27026	    443408 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8       	  147583	     79870 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8            	13974954	       863.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                   	 1000000	     11730 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                   	  639312	     18991 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                          	 1482436	      8049 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8             	 1450267	      8357 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8           	129175284	        91.68 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8   	73754100	       174.9 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                     	   14276	    825779 ns/op	  916973 B/op	    6011 allocs/op
Benchmark_Unmarshal_JsonvalueNoCopy-8                    	   16281	    759251 ns/op	  851435 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                          	 2654900	      4439 ns/op	    2224 B/op	       5 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	488.685s
```
