# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=2s && go version
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  294295	      7938 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  466478	      5062 ns/op	    1617 B/op	      21 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  697129	      3492 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  379947	      6036 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_interface_json-8                        	  213194	     11340 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_interface_sonic-8                       	  445749	      5430 ns/op	    5011 B/op	      51 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter-8                    	  286650	      8451 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_interface_json_blob-8                   	    2504	    990808 ns/op	  414635 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter_blob-8               	    2790	   1473645 ns/op	  457071 B/op	   13302 allocs/op
Benchmark_Unmarshal_interface_json-8                            	  246636	      9870 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_interface_sonic-8                           	  472128	      5412 ns/op	    4918 B/op	      51 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  225663	     11909 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  134397	     17612 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  316526	      7703 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_latest-8                          	  295659	      8113 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  301324	      8060 ns/op	    9648 B/op	      59 allocs/op
Benchmark__Marshal__map_interface_json-8                        	  178966	     13566 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  455828	      5252 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  451641	      5561 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1000000	      2160 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   66747	     37236 ns/op	   26352 B/op	     349 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	   79520	     30551 ns/op	   23040 B/op	     342 allocs/op
Benchmark__Import___结构体_jsonvalue_beta-8                        	   77769	     32856 ns/op	   23040 B/op	     342 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  170487	     13909 ns/op	   10902 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  169934	     13861 ns/op	   15978 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	  184538	     14223 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	 1000000	      2371 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   55200	     41864 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  384169	      5999 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  291406	      8073 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  543280	      4495 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   83427	     26672 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  767538	      3168 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 2813214	       839.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   29110	     82785 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 2646948	       851.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  266882	      9122 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  288812	      8328 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    2836	    863294 ns/op	  977068 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	29455364	        80.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	16394924	       145.2 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2816	    871894 ns/op	  976997 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    3217	    830974 ns/op	  911463 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  511050	      4484 ns/op	    3312 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   52806	     45613 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   13869	    168466 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6277	    383470 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3552	    703403 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2080	   1127641 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1581	   1515289 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   13821	    164680 ns/op	   38572 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4615	    519833 ns/op	   96352 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    2113	   1069588 ns/op	  173336 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1284	   1783813 ns/op	  269528 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     879	   2662526 ns/op	  384926 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     632	   3807361 ns/op	  519532 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2767	    871396 ns/op	  977057 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2763	    876718 ns/op	  977066 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2734	    896607 ns/op	  977074 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2646	    884687 ns/op	  977072 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2584	    902536 ns/op	  977074 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2690	    904053 ns/op	  977066 B/op	    5813 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test	166.119s
go version go1.19 darwin/amd64

```
