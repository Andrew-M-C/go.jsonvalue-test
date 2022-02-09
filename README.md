# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=1s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJson_Struct-8                          	  142930	      8235 ns/op	    1112 B/op	      25 allocs/op
Benchmark__Marshal__GoStdJson_Struct-8                          	  207270	      5996 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_GoStdJson_MapItf_blob-8                     	    1216	   1008316 ns/op	  414617 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJson_MapItf-8                          	  102210	     11908 ns/op	    4480 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJson_MapItf-8                          	   79012	     14943 ns/op	    6185 B/op	     121 allocs/op
Benchmark_Unmarshal_Sonic_Struct-8                              	  232362	      5023 ns/op	    1611 B/op	      21 allocs/op
Benchmark__Marshal__Sonic_Struct-8                              	  197806	      5814 ns/op	    2285 B/op	       9 allocs/op
Benchmark_Unmarshal_Sonic_MapItf_blob-8                         	    2378	    518482 ns/op	  471410 B/op	    4913 allocs/op
Benchmark_Unmarshal_Sonic_MapItf-8                              	  230642	      5224 ns/op	    4974 B/op	      51 allocs/op
Benchmark__Marshal__Sonic_MapItf-8                              	  341214	      3437 ns/op	     883 B/op	       4 allocs/op
Benchmark_Unmarshal_Jsoniter_Struct-8                           	  193605	      6077 ns/op	    1720 B/op	      56 allocs/op
Benchmark__Marshal__Jsoniter_Struct-8                           	  206613	      5870 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_Jsoniter_MapItf_blob-8                      	    1333	    887070 ns/op	  457073 B/op	   13302 allocs/op
Benchmark_Unmarshal_Jsoniter_MapItf-8                           	  140618	      8931 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	   81470	     15017 ns/op	    6185 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	  513456	      2313 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   28728	     41835 ns/op	   12659 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  194631	      6058 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  140308	      8445 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  270566	      4652 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   43350	     27997 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Easyjson_Struct-8                           	  347764	      3409 ns/op	     784 B/op	      19 allocs/op
Benchmark__Marshal__Easyjson_Struct-8                           	  584731	      2175 ns/op	    1240 B/op	       5 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  378313	      3221 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 1561994	       769.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   15802	     75584 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 1581488	       751.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  147976	      8065 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  103242	     11265 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	   63919	     18588 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  148083	      8152 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  152042	      8100 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  134790	      8466 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    1408	    870399 ns/op	  917116 B/op	    6012 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	14725621	        80.79 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	 7971552	       153.3 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    1287	    935601 ns/op	  916976 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    1461	    759197 ns/op	  851435 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  249309	      4783 ns/op	    2368 B/op	       6 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   29379	     40960 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	    7580	    156714 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    3412	    345994 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    1930	    613125 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    1273	    960925 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	     867	   1378488 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_10_percent-8        	    5794	    194561 ns/op	  761394 B/op	      10 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_20_percent-8        	    2383	    491311 ns/op	 1506567 B/op	      21 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_30_percent-8        	    1306	    883229 ns/op	 2240915 B/op	      32 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_40_percent-8        	     867	   1372565 ns/op	 2972644 B/op	      43 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_50_percent-8        	     624	   2022342 ns/op	 3678427 B/op	      54 allocs/op
Benchmark____Get____Sonic_ReadOneChain_Blob_60_percent-8        	     414	   2830130 ns/op	 4416181 B/op	      65 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	    7308	    169587 ns/op	   38572 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    2280	    525356 ns/op	   96351 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1089	   1071753 ns/op	  173335 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	     663	   1813107 ns/op	  269526 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     410	   2842988 ns/op	  384921 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     300	   3853187 ns/op	  519526 B/op	   67965 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark	79.254s

```
