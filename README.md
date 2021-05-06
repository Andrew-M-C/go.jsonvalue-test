# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=2s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJsonStruct-8                           	  254887	      8775 ns/op	    1144 B/op	      25 allocs/op
Benchmark__Marshal__GoStdJsonStruct-8                           	  353425	      6859 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf_blob-8                      	    2150	   1121728 ns/op	  414657 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf-8                           	  191503	     13040 ns/op	    4512 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJsonMapItf-8                           	  139951	     17140 ns/op	    5865 B/op	     121 allocs/op
Benchmark_Unmarshal_JsoniterStruct-8                            	  357116	      6890 ns/op	    1720 B/op	      56 allocs/op
Benchmark__Marshal__JsoniterStruct-8                            	  343639	      6843 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_EasyjsonStruct-8                            	  629749	      4017 ns/op	     784 B/op	      19 allocs/op
Benchmark__Marshal__EasyjsonStruct-8                            	 1000000	      2463 ns/op	    1240 B/op	       5 allocs/op
Benchmark_Unmarshal_JsoniterMapItf_blob-8                       	    2528	    966703 ns/op	  457049 B/op	   13302 allocs/op
Benchmark_Unmarshal_JsoniterMapItf-8                            	  255297	      9442 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__JsoniterMapItf-8                            	  127723	     17132 ns/op	    5865 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	  931263	      2461 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   50624	     45237 ns/op	   12659 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  361866	      6568 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  265844	      9118 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  497700	      4936 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   82556	     29967 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  671830	      3705 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 2743628	       887.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   29400	     82223 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 2740824	       917.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  273925	      8826 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  203238	     11795 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  125674	     18927 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  310831	      7684 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  303901	      7928 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    3092	    799450 ns/op	  917030 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	26012683	        92.07 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	13932954	       164.6 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2976	    785346 ns/op	  916975 B/op	    6011 allocs/op
Benchmark_Unmarshal_JsonvalueNoCopy-8                           	    3156	    716319 ns/op	  851432 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  546277	      4376 ns/op	    2224 B/op	       5 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   53822	     44646 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   14209	    169907 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6487	    373631 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3654	    658874 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2349	   1026668 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1636	   1465678 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   13387	    178256 ns/op	   38571 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4318	    565711 ns/op	   96348 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    2090	   1172316 ns/op	  173330 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1203	   1926726 ns/op	  269519 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     825	   2927651 ns/op	  384913 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     580	   4099281 ns/op	  519511 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2816	    769587 ns/op	  917038 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    3142	    771061 ns/op	  917042 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2953	    796379 ns/op	  917033 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2979	    792485 ns/op	  917024 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    3064	    798794 ns/op	  917036 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    3073	    802965 ns/op	  917033 B/op	    6011 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	136.637s
```
