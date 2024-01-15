# Test Result 

```
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  269496	      8094 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  466189	      4960 ns/op	    1527 B/op	      22 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  718216	      3368 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_manual_by_jsonvalue-8                   	  172094	     13129 ns/op	   17104 B/op	     117 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  332359	      6068 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_any_json-8                              	  212064	     11174 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_any_sonic-8                             	  465806	      4892 ns/op	    4588 B/op	      52 allocs/op
Benchmark_Unmarshal_map_any_jsoniter-8                          	  278262	      8456 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_any_json_blob-8                         	    2377	   1004319 ns/op	  423730 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_any_jsoniter_blob-8                     	    2745	    895234 ns/op	  466212 B/op	   13302 allocs/op
Benchmark_Unmarshal_map_any_sonic_blob-8                        	    4989	    539037 ns/op	  473274 B/op	    4914 allocs/op
Benchmark_Unmarshal_any_json-8                                  	  228176	     10127 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_any_sonic-8                                 	  480010	      5654 ns/op	    4592 B/op	      52 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  223201	     10907 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  139278	     18421 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  284257	      8272 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_3_3-8                          	  291470	      8309 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_latest-8                          	  325806	     11492 ns/op	    9616 B/op	      32 allocs/op
Benchmark__Marshal__map_any_json-8                              	  186712	     12798 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  404860	      7156 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_sonic-8                                 	  510144	      4472 ns/op	    2025 B/op	       9 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  968521	      2485 ns/op	    1889 B/op	       7 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1000000	      2010 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   85227	     28129 ns/op	   17592 B/op	     304 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	  103578	     23018 ns/op	   14280 B/op	     297 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  159615	     18306 ns/op	   10906 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  149827	     14498 ns/op	   12939 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	  168151	     19010 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	  896138	      3797 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   55968	     42543 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  367641	      6238 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  274466	      8682 ns/op	    3025 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  491296	      4596 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   84448	     28907 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  783763	      3190 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 3053454	       931.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   31581	    110247 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 1000000	      2583 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  281928	      9298 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  273660	      9207 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    2511	    938705 ns/op	  986183 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	27330741	        87.34 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	14535043	       171.2 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2602	    928590 ns/op	  986115 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    2989	    742392 ns/op	  920557 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  445809	      4960 ns/op	    3312 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   55009	     42762 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   14439	    164117 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6530	    373026 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3751	    694900 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2370	   1016452 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1624	   1448083 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   12825	    183607 ns/op	   38574 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4096	    595140 ns/op	   96355 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1885	   1211171 ns/op	  173344 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1117	   2250863 ns/op	  269538 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     795	   3230216 ns/op	  384939 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     475	   4491361 ns/op	  519549 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2545	    925888 ns/op	  978084 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2416	    978325 ns/op	  978074 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2596	    835150 ns/op	  978083 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2948	    824605 ns/op	  978069 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2653	    911031 ns/op	  978076 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2788	    865647 ns/op	  978069 B/op	    2914 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark	190.876s
go version go1.21.0 darwin/amd64
```
