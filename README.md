goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  278172	      8482 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  436108	      5337 ns/op	    1497 B/op	      21 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  654159	      3710 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  364047	      6334 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_any_json-8                              	  205944	     11869 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_any_sonic-8                             	  476872	      5000 ns/op	    4540 B/op	      51 allocs/op
Benchmark_Unmarshal_map_any_jsoniter-8                          	  267441	      9266 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_any_json_blob-8                         	    2260	   1039231 ns/op	  414619 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_any_jsoniter_blob-8                     	    2673	    912458 ns/op	  457071 B/op	   13302 allocs/op
Benchmark_Unmarshal_any_json-8                                  	  237934	     10270 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_any_sonic-8                                 	  458076	      5057 ns/op	    4551 B/op	      51 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  216398	     11033 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  132889	     18266 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  311277	      7869 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_latest-8                          	  291501	      8390 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  286314	      8058 ns/op	    9648 B/op	      59 allocs/op
Benchmark__Marshal__map_any_json-8                              	  167246	     13568 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  420160	      5581 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  418749	      5574 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1000000	      2106 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   79479	     30476 ns/op	   18120 B/op	     337 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	   94078	     25049 ns/op	   14808 B/op	     330 allocs/op
Benchmark__Import___结构体_jsonvalue_beta-8                        	   95097	     25295 ns/op	   14808 B/op	     330 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  168106	     14370 ns/op	   10903 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  161674	     14667 ns/op	   14654 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	  171906	     13538 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	  988339	      2425 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   56395	     42936 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  380377	      6279 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  274083	      8652 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  513423	      4956 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   83995	     28696 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  808168	      2874 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 3352128	       721.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   33168	     70600 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 2964844	       811.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  291436	      8235 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  276459	      9288 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    2822	    886245 ns/op	  977057 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	28911877	        81.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	15397184	       152.1 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2732	    890690 ns/op	  976993 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    3175	    734432 ns/op	  911446 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  491637	      4670 ns/op	    3312 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   60729	     40178 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   15499	    160498 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6864	    338058 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3964	    605535 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2497	    943757 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1800	   1340332 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   13024	    179988 ns/op	   38572 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4324	    560262 ns/op	   96352 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    2140	   1134985 ns/op	  173337 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1245	   1923878 ns/op	  269528 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     820	   2901671 ns/op	  384928 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     574	   4106698 ns/op	  519533 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2674	    891555 ns/op	  977079 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2695	    893055 ns/op	  977070 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2641	    901437 ns/op	  977078 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2572	   1061364 ns/op	  977055 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2553	    925623 ns/op	  977066 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2527	    931654 ns/op	  977060 B/op	    5813 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test	163.479s
go version go1.19 darwin/amd64

```
