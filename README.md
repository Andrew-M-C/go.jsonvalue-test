# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=2s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJsonStruct-8                           	  283948	      8335 ns/op	    1112 B/op	      25 allocs/op
Benchmark__Marshal__GoStdJsonStruct-8                           	  468324	      5452 ns/op	    1881 B/op	       6 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf_blob-8                      	    2412	   1032342 ns/op	  414644 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf-8                           	  198566	     12189 ns/op	    4480 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJsonMapItf-8                           	  156650	     15273 ns/op	    6185 B/op	     121 allocs/op
Benchmark_Unmarshal_JsoniterStruct-8                            	  375802	      6413 ns/op	    1720 B/op	      56 allocs/op
Benchmark__Marshal__JsoniterStruct-8                            	  423270	      5670 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_EasyjsonStruct-8                            	  675633	      3562 ns/op	     784 B/op	      19 allocs/op
Benchmark__Marshal__EasyjsonStruct-8                            	 1000000	      2329 ns/op	    1240 B/op	       5 allocs/op
Benchmark_Unmarshal_JsoniterMapItf_blob-8                       	    2622	    925246 ns/op	  457028 B/op	   13302 allocs/op
Benchmark_Unmarshal_JsoniterMapItf-8                            	  269809	      8962 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__JsoniterMapItf-8                            	  160050	     15124 ns/op	    6185 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	 1000000	      2389 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   55204	     43972 ns/op	   12659 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  372058	      6296 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  280410	      8912 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  502022	      4654 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   85657	     28709 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  750291	      3252 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 3015685	       788.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   31124	     76819 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 3053313	      1158 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  275320	      8282 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  209709	     11430 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  129127	     18738 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  314898	      7831 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  299674	      8116 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  288716	      8245 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    3079	    820548 ns/op	  917030 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	28635333	        81.69 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	15328794	       155.6 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2709	    808581 ns/op	  916982 B/op	    6011 allocs/op
Benchmark_Unmarshal_JsonvalueNoCopy-8                           	    2701	    794418 ns/op	  851429 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  507854	      4671 ns/op	    2368 B/op	       6 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   57507	     41875 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   14925	    160134 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6692	    353263 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3843	    625515 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2439	    981672 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1725	   1407233 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   13341	    173294 ns/op	   38571 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4324	    552389 ns/op	   96348 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1981	   1138117 ns/op	  173331 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1260	   1959927 ns/op	  269519 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     840	   2834055 ns/op	  384913 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     608	   4071067 ns/op	  519512 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2745	    808487 ns/op	  917034 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2901	    817095 ns/op	  917040 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2611	    904495 ns/op	  917042 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2654	    904284 ns/op	  917033 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2592	    933547 ns/op	  917035 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2559	    921621 ns/op	  917039 B/op	    6011 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	139.672s
```
