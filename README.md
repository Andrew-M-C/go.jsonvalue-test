# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test resule for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=10s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJsonStruct-8                    	  993850	     11745 ns/op	    1520 B/op	      49 allocs/op
Benchmark__Marshal__GoStdJsonStruct-8                    	 3571764	      3292 ns/op	     640 B/op	       1 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf_blob-8               	   10000	   1033259 ns/op	  414658 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf-8                    	  950488	     12124 ns/op	    4512 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJsonMapItf-8                    	  746385	     16226 ns/op	    5865 B/op	     121 allocs/op
Benchmark_Unmarshal_JsoniterStruct-8                     	 4240194	      2812 ns/op	     768 B/op	      22 allocs/op
Benchmark__Marshal__JsoniterStruct-8                     	 3641932	      3271 ns/op	     640 B/op	       1 allocs/op
Benchmark_Unmarshal_JsoniterMapItf_blob-8                	   13083	    911518 ns/op	  457041 B/op	   13302 allocs/op
Benchmark_Unmarshal_JsoniterMapItf-8                     	 1313599	      8989 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__JsoniterMapItf-8                     	  738814	     16397 ns/op	    5865 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                           	 4907690	      2439 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8         	 1835197	      6345 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8              	 1362438	      8726 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8              	  934304	     12857 ns/op	    3665 B/op	     197 allocs/op
Benchmark____Get____Jsoniter_blob_ReadOneChain-8         	   28911	    413569 ns/op	  104263 B/op	    6788 allocs/op
Benchmark_Unmarshal_Jsonparser_Full_Blob-8               	   32892	    356057 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8       	  202369	     59646 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8            	17395621	       690.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                   	 1000000	     11078 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                   	  672234	     18234 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                          	 1551367	      7650 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8             	 1513297	      7865 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8           	131101486	        91.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8   	73975131	       161.2 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                     	   15094	    787331 ns/op	  916970 B/op	    6011 allocs/op
Benchmark_Unmarshal_JsonvalueNoCopy-8                    	   17750	    681302 ns/op	  851433 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                          	 2691966	      4429 ns/op	    2224 B/op	       5 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	433.393s
```
