# Test Result 

```
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  278119	      7815 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  480157	      4928 ns/op	    1496 B/op	      21 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  712092	      3342 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  401158	      5839 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_any_json-8                              	  217786	     11021 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_any_sonic-8                             	  436867	      4792 ns/op	    4542 B/op	      51 allocs/op
Benchmark_Unmarshal_map_any_jsoniter-8                          	  265198	      8313 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_any_json_blob-8                         	    2497	    960434 ns/op	  414635 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_any_jsoniter_blob-8                     	    2784	    847731 ns/op	  457070 B/op	   13302 allocs/op
Benchmark_Unmarshal_map_any_sonic_blob-8                        	    4598	    492926 ns/op	  463989 B/op	    4913 allocs/op
Benchmark_Unmarshal_any_json-8                                  	  252283	      9580 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_any_sonic-8                                 	  508942	      4700 ns/op	    4550 B/op	      51 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  230788	     10405 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  137710	     17359 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  318654	      7435 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_latest-8                          	  306370	      7911 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  321985	      7356 ns/op	    9616 B/op	      32 allocs/op
Benchmark__Marshal__map_any_json-8                              	  190083	     12651 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  458941	      5206 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_sonic-8                                 	  589070	      4102 ns/op	    2065 B/op	       9 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  457198	      5219 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1210135	      1973 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   82521	     28582 ns/op	   18120 B/op	     337 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	  101342	     23386 ns/op	   14808 B/op	     330 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  177090	     13631 ns/op	   10902 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  178653	     13805 ns/op	   14599 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	  188378	     12632 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	 1000000	      2156 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   60922	     39552 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  408208	      5718 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  302323	      7889 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  549182	      4305 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   91376	     26075 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  759166	      3141 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 2900695	       822.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   31236	     75920 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 2923836	       830.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  282600	      8418 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  297517	      8126 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    2816	    846009 ns/op	  977063 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	30817838	        77.65 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	16778632	       160.9 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    1785	   1508291 ns/op	  977004 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop_blob-8                    	    3063	    757261 ns/op	  968895 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    2937	    689185 ns/op	  911456 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  489666	      4490 ns/op	    3312 B/op	       7 allocs/op
Benchmark__Marshal__Jsonvalue_develop-8                         	  529689	      4412 ns/op	    3312 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   55303	     42797 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   14571	    163528 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6667	    360132 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    3728	    628600 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2456	    989869 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1692	   1402149 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   14956	    159947 ns/op	   38573 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4629	    504253 ns/op	   96351 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    2289	   1031781 ns/op	  173337 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1370	   1750581 ns/op	  269531 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     900	   2629224 ns/op	  384928 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     644	   3747333 ns/op	  519536 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2739	    868091 ns/op	  977085 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2786	    870714 ns/op	  977067 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2782	    872094 ns/op	  977080 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    2794	    872923 ns/op	  977069 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2746	    880189 ns/op	  977055 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2694	    885687 ns/op	  977070 B/op	    5813 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test	169.419s
go version go1.19 darwin/amd64

```
