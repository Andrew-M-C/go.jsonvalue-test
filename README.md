# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=1s
goos: darwin
goarch: amd64
pkg: github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_结构体_json-8                                  	  142317	      8795 ns/op	    1112 B/op	      25 allocs/op
Benchmark_Unmarshal_结构体_sonic-8                                 	  197630	      6411 ns/op	    1632 B/op	      21 allocs/op
Benchmark_Unmarshal_结构体_easyjson-8                              	  319024	      3624 ns/op	     784 B/op	      19 allocs/op
Benchmark_Unmarshal_结构体_jsoniter-8                              	  200072	      6103 ns/op	    1720 B/op	      56 allocs/op
Benchmark_Unmarshal_map_interface_json-8                        	  106363	     11200 ns/op	    4480 B/op	     128 allocs/op
Benchmark_Unmarshal_map_interface_sonic-8                       	  222970	      5234 ns/op	    4997 B/op	      51 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter-8                    	  141306	      8250 ns/op	    4521 B/op	     136 allocs/op
Benchmark_Unmarshal_map_interface_json_blob-8                   	    1225	    965584 ns/op	  414652 B/op	   11408 allocs/op
Benchmark_Unmarshal_map_interface_jsoniter_blob-8               	    1392	    847972 ns/op	  457076 B/op	   13302 allocs/op
Benchmark_Unmarshal_interface_json-8                            	  123634	     10021 ns/op	    4272 B/op	     117 allocs/op
Benchmark_Unmarshal_interface_sonic-8                           	  217779	      5371 ns/op	    4885 B/op	      51 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  108270	     11157 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	   68006	     19081 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_1_1-8                          	  154202	      7691 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  150699	      7558 ns/op	    8320 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_develop-8                         	  162558	      7436 ns/op	    8320 B/op	      59 allocs/op
Benchmark__Marshal__map_interface_json-8                        	   83882	     14435 ns/op	    6185 B/op	     121 allocs/op
Benchmark__Marshal__结构体_json-8                                  	  201583	      5723 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_jsoniter-8                              	  207591	      5753 ns/op	    1882 B/op	       6 allocs/op
Benchmark__Marshal__结构体_easyjson-8                              	  521811	      2366 ns/op	    1240 B/op	       5 allocs/op
Benchmark__Marshal__结构体_jsonvalue-8                             	   27642	     39982 ns/op	   23984 B/op	     348 allocs/op
Benchmark__Import___结构体_jsonvalue-8                             	   38191	     29660 ns/op	   20432 B/op	     341 allocs/op
Benchmark__Import___结构体_jsonvalue_beta-8                        	   41816	     28681 ns/op	   20432 B/op	     341 allocs/op
Benchmark__Import___结构体_jsonvalue_json中转-8                      	   85194	     13749 ns/op	    9571 B/op	      64 allocs/op
Benchmark__Import___结构体_jsonvalue_sonic中转-8                     	   90510	     12715 ns/op	   13660 B/op	      67 allocs/op
Benchmark__Marshal__Jsoniter_MapItf-8                           	   82110	     14334 ns/op	    6185 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	  566554	      2239 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   28827	     42824 ns/op	   12659 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  199940	      5980 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  145426	      8243 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  273231	      4454 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   44958	     26569 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  389695	      3137 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 1626951	       725.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   17390	     69807 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 1651392	       711.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  150801	      7712 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  153502	      7734 ns/op	    8320 B/op	      59 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    1478	    775468 ns/op	  841855 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	15121879	        77.69 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	 8302702	       145.7 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    1393	    812793 ns/op	  841743 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_NoCopy-8                          	    1612	    686008 ns/op	  776201 B/op	    5810 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  256501	      4690 ns/op	    3552 B/op	       7 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   29874	     40004 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	    7876	    154841 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    3480	    360041 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    1622	    650095 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    1190	    948552 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	     904	   1319657 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	    7131	    165898 ns/op	   38571 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    2282	    520144 ns/op	   96350 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    1102	   1071902 ns/op	  173334 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	     664	   2074838 ns/op	  269524 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     435	   2750043 ns/op	  384921 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     316	   3878384 ns/op	  519521 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    1478	    802965 ns/op	  841864 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    1488	    809192 ns/op	  841863 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    1434	    807907 ns/op	  841862 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    1441	    824380 ns/op	  841865 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    1436	    812454 ns/op	  841868 B/op	    5811 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    1479	    815053 ns/op	  841852 B/op	    5811 allocs/op
PASS
ok  	github.com/Andrew-M-C/go.jsonvalue-test/benchmark_test	83.698s
```
