# Test Result 

```
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  292036	      8361 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  438756	      5424 ns/op	    1529 B/op	      22 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  662464	      3443 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_manual_by_jsonvalue-8                   	  191065	     12538 ns/op	   17104 B/op	     117 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  392132	      6269 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_any_json-8                              	  200018	     12034 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_any_sonic-8                             	  476901	      4819 ns/op	    4581 B/op	      52 allocs/op
Benchmark_Unmarshal_map_any_jsoniter-8                          	  274323	      8616 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_any_json_blob-8                         	    2464	   1260515 ns/op	  414640 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_any_jsoniter_blob-8                     	    2732	    896598 ns/op	  457070 B/op	   13302 allocs/op
Benchmark_Unmarshal_map_any_sonic_blob-8                        	    4368	    484602 ns/op	  464753 B/op	    4914 allocs/op
Benchmark_Unmarshal_any_json-8                                  	  241207	      9937 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_any_sonic-8                                 	  515865	      4652 ns/op	    4591 B/op	      52 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  244132	     10217 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  144428	     16698 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  308692	      7545 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_3_3-8                          	  283809	      7818 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  317084	      7077 ns/op	    9616 B/op	      32 allocs/op
Benchmark__Marshal__map_any_json-8                              	  176785	     13479 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  443728	      5313 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_sonic-8                                 	  611104	      4049 ns/op	    1996 B/op	       9 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  447393	      5328 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1000000	      2111 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   71433	     29075 ns/op	   18120 B/op	     337 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	   98504	     24657 ns/op	   14808 B/op	     330 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  174692	     13968 ns/op	   10902 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  178287	     29050 ns/op	   12066 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	  182209	     13264 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	  985748	      2252 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   55322	     42247 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  391555	      5986 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  284059	     10606 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  511243	      4815 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   80816	     31605 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  735368	      3150 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 2400475	       878.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   33997	     71522 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 2906630	       806.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  276318	      8098 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  300937	      8108 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    2832	    900733 ns/op	  977059 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	23218699	        91.17 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	15254594	       168.0 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    2805	    848326 ns/op	  977019 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop_blob-8                    	    3322	    749761 ns/op	  968902 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    3051	    866160 ns/op	  911449 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  500131	      4524 ns/op	    3312 B/op	       7 allocs/op
Benchmark__Marshal__Jsonvalue_develop-8                         	  515959	      4748 ns/op	    5400 B/op	       3 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   60000	     40010 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   15648	    153822 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    6961	    344065 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    4033	    640570 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2581	    935022 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1798	   1335834 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   13554	    174574 ns/op	   38572 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4138	    628498 ns/op	   96351 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1874	   1203121 ns/op	  173336 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1100	   2399234 ns/op	  269527 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     781	   3483911 ns/op	  384924 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     549	   4547839 ns/op	  519528 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    3078	    766447 ns/op	  968973 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    3273	    758746 ns/op	  968962 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    3126	    835099 ns/op	  968953 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    3080	    780475 ns/op	  968947 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    3036	    776272 ns/op	  968968 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2572	    821112 ns/op	  968966 B/op	    2914 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark	180.845s
go version go1.20.5 darwin/amd64


```
