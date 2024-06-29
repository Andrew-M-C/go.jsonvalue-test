# Test Result

```
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  250555	      9214 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  398770	      5289 ns/op	    1530 B/op	      22 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  608262	      3886 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_manual_by_jsonvalue-8                   	  162153	     13542 ns/op	   17104 B/op	     117 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  347964	      6879 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_any_json-8                              	  200397	     11868 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_any_sonic-8                             	  468549	      5112 ns/op	    4585 B/op	      52 allocs/op
Benchmark_Unmarshal_map_any_jsoniter-8                          	  243999	      9276 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_any_json_blob-8                         	    2276	   1134169 ns/op	  423750 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_any_jsoniter_blob-8                     	    2521	    979129 ns/op	  466204 B/op	   13302 allocs/op
Benchmark_Unmarshal_map_any_sonic_blob-8                        	    4318	    552069 ns/op	  473236 B/op	    4914 allocs/op
Benchmark_Unmarshal_any_json-8                                  	  224694	     10527 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_any_sonic-8                                 	  486903	      5179 ns/op	    4597 B/op	      52 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  228272	     11628 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  117376	     23236 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  195751	     12694 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_3_3-8                          	  187226	     15525 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_latest-8                          	  186278	     14100 ns/op	    9616 B/op	      32 allocs/op
Benchmark__Marshal__map_any_json-8                              	  103904	     21894 ns/op	    5033 B/op	     103 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  230912	      9144 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__map_any_sonic-8                             	  527053	      4224 ns/op	     725 B/op	       4 allocs/op
Benchmark__Marshal__结构体_sonic-8                                 	  475858	      4877 ns/op	    2020 B/op	       9 allocs/op
Benchmark__Marshal__map_any_jsoniter-8                          	  158620	     14587 ns/op	    5440 B/op	      99 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  729784	      3042 ns/op	    1889 B/op	       7 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	 1000000	      2467 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   73206	     33730 ns/op	   17592 B/op	     304 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	   92600	     30474 ns/op	   14280 B/op	     297 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	  126811	     22908 ns/op	   10906 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	  107001	     28454 ns/op	   12788 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	   85306	     30917 ns/op	    5033 B/op	     103 allocs/op
Benchmark____Get____Jsoniter-8                                  	  502749	      5397 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   25704	     82586 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  218757	     10115 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  150880	     15051 ns/op	    3025 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  298924	     12457 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   49621	     45790 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  338791	      7321 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 1000000	      2800 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   15301	    179305 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 1706084	      1379 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  180036	     13123 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  172562	     14245 ns/op	    9648 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    1948	   1403712 ns/op	  986213 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	17200914	       119.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	11371784	       246.9 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    1587	   1652673 ns/op	  986093 B/op	    5813 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    2023	   1030440 ns/op	  920566 B/op	    5812 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  311330	      6954 ns/op	    3312 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   37837	     77486 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	    9382	    283467 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    3294	    636244 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    2230	   1030772 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    1693	   1334216 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    1234	   1879788 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	    9735	    265138 ns/op	   38574 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    3396	    714044 ns/op	   96357 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1788	   1461323 ns/op	  173345 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1017	   2420647 ns/op	  269541 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     698	   3582337 ns/op	  384943 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     465	   4964821 ns/op	  519553 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2258	   1137336 ns/op	  978101 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    2023	   1136447 ns/op	  978114 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2076	   1245670 ns/op	  978100 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    1978	   1531504 ns/op	  978121 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    1472	   1452272 ns/op	  978143 B/op	    2914 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    1580	   1540789 ns/op	  978151 B/op	    2914 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark	198.894s
go version go1.21.0 darwin/amd64
```
