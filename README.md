# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=1s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  144528	      8030 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  240597	      4948 ns/op	    1620 B/op	      21 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  362871	      3356 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  201948	      6131 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_interface_json-8                        	  104917	     11478 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_interface_sonic-8                       	  229995	      5137 ns/op	    4941 B/op	      51 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter-8                    	  140928	      8458 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_interface_json_blob-8                   	    1178	    985763 ns/op	  414613 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter_blob-8               	    1359	    878887 ns/op	  457062 B/op	   13302 allocs/op
Benchmark_Unmarshal_interface_json-8                            	  121664	      9808 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_interface_sonic-8                           	  228894	      4984 ns/op	    4861 B/op	      51 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  110179	     10726 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	   71094	     17717 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  154677	      7770 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  156442	      7663 ns/op	    8320 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  153625	      7726 ns/op	    8320 B/op	      59 allocs/op
Benchmark__Marshal__map_interface_json-8                        	   82711	     14552 ns/op	    6185 B/op	     121 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  206121	      5786 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  206034	      5770 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	  591543	      2131 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	   82446	     14606 ns/op	    6185 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	  538974	      2276 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   28891	     41248 ns/op	   12660 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  197341	      6061 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  141849	      8334 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  267992	      4500 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   44136	     27787 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  395827	      3054 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 1666539	       722.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   17634	     67990 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 1620021	       742.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  165807	      7349 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  148784	      7844 ns/op	    8320 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    1557	    781694 ns/op	  841870 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	15232558	        78.41 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	 8077336	       151.4 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    1604	    774243 ns/op	  841742 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    1594	    698940 ns/op	  776204 B/op	    5810 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  263047	      4531 ns/op	    3552 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   32016	     37605 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	    8557	    140447 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    3912	    310888 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    2182	    547787 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    1408	    852751 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	     963	   1231516 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	    7024	    169237 ns/op	   38572 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    2208	    530854 ns/op	   96351 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1107	   1089681 ns/op	  173334 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	     649	   1849233 ns/op	  269526 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     434	   2845511 ns/op	  384923 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     292	   4037834 ns/op	  519528 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    1533	    777603 ns/op	  841863 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    1548	    784982 ns/op	  841855 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    1502	    781319 ns/op	  841856 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    1533	    794562 ns/op	  841867 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    1486	    789098 ns/op	  841849 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    1524	    796229 ns/op	  841857 B/op	    5811 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	76.885s
```
