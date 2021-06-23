# go.jsonvalue-test

A repo for go.jsonvalue pprof and benchmark test.

# Test result for reference: 

```
% go test -bench=. -run=none -benchmem -benchtime=2s
goos: darwin
goarch: amd64
pkg: go.jsonvalue-test/benchmark_test
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
Benchmark_Unmarshal_GoStdJsonStruct-8                           	  269788	      9065 ns/op	    1144 B/op	      25 allocs/op
Benchmark__Marshal__GoStdJsonStruct-8                           	  352569	      7034 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf_blob-8                      	    2196	   1054230 ns/op	  414668 B/op	   11408 allocs/op
Benchmark_Unmarshal_GoStdJsonMapItf-8                           	  189753	     12304 ns/op	    4512 B/op	     128 allocs/op
Benchmark__Marshal__GoStdJsonMapItf-8                           	  150313	     16007 ns/op	    5865 B/op	     121 allocs/op
Benchmark_Unmarshal_JsoniterStruct-8                            	  361234	      6654 ns/op	    1720 B/op	      56 allocs/op
Benchmark__Marshal__JsoniterStruct-8                            	  351315	      6689 ns/op	    1882 B/op	       6 allocs/op
Benchmark_Unmarshal_EasyjsonStruct-8                            	  652816	      3846 ns/op	     784 B/op	      19 allocs/op
Benchmark__Marshal__EasyjsonStruct-8                            	 1000000	      2386 ns/op	    1240 B/op	       5 allocs/op
Benchmark_Unmarshal_JsoniterMapItf_blob-8                       	    2664	    899312 ns/op	  457055 B/op	   13302 allocs/op
Benchmark_Unmarshal_JsoniterMapItf-8                            	  255264	      9167 ns/op	    4521 B/op	     136 allocs/op
Benchmark__Marshal__JsoniterMapItf-8                            	  147895	     15881 ns/op	    5865 B/op	     121 allocs/op
Benchmark____Get____Jsoniter-8                                  	  986508	      2416 ns/op	     912 B/op	      36 allocs/op
Benchmark____Get____Jsoniter_Full-8                             	   54272	     44164 ns/op	   12659 B/op	     671 allocs/op
Benchmark____Get____Jsoniter_AndGetParsedValue-8                	  361194	      6556 ns/op	    2112 B/op	     103 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain-8                     	  250202	      9041 ns/op	    3024 B/op	     139 allocs/op
Benchmark____Get____Jsoniter_ReadLevelOne-8                     	  494145	      4957 ns/op	    1344 B/op	      73 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob-8                	   79394	     31576 ns/op	    4913 B/op	     469 allocs/op
Benchmark_Unmarshal_Jsonparser_Full-8                           	  810127	      3098 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne-8                   	 3574604	       720.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadLevelOne_Blob-8              	   39124	     65601 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain-8                   	 3148964	       706.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob-8              	  361150	      6847 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_3-8                          	  211962	     10971 ns/op	    7184 B/op	     104 allocs/op
Benchmark_Unmarshal_Jsonvalue_v1_0_4-8                          	  131930	     17921 ns/op	   15824 B/op	     188 allocs/op
Benchmark_Unmarshal_Jsonvalue-8                                 	  309998	      7562 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain-8                    	  313290	      7742 ns/op	    9072 B/op	      61 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob-8               	    3146	    780362 ns/op	  917019 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterface-8                  	25846873	        92.60 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonvalue_GetByInterfaceCaseless-8          	15059178	       165.5 ns/op	      72 B/op	       2 allocs/op
Benchmark_Unmarshal_Jsonvalue_blob-8                            	    3153	    777072 ns/op	  916979 B/op	    6011 allocs/op
Benchmark_Unmarshal_JsonvalueNoCopy-8                           	    3106	    719167 ns/op	  851433 B/op	    6010 allocs/op
Benchmark__Marshal__Jsonvalue-8                                 	  547471	      4357 ns/op	    2224 B/op	       5 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_10_percent-8   	   70198	     65496 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_20_percent-8   	   16776	    129998 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_30_percent-8   	    8422	    307469 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_40_percent-8   	    4786	    504724 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_50_percent-8   	    2923	    874240 ns/op	       0 B/op	       0 allocs/op
Benchmark_Unmarshal_Jsonparser_ReadOneChain_Blob_60_percent-8   	    2109	   1152100 ns/op	       0 B/op	       0 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_10_percent-8     	   12747	    186608 ns/op	   38571 B/op	    2865 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_20_percent-8     	    4045	    579362 ns/op	   96348 B/op	    9085 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_30_percent-8     	    2018	   1190685 ns/op	  173331 B/op	   18705 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_40_percent-8     	    1180	   2029144 ns/op	  269519 B/op	   31725 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_50_percent-8     	     783	   3037155 ns/op	  384913 B/op	   48145 allocs/op
Benchmark____Get____Jsoniter_ReadOneChain_Blob_60_percent-8     	     562	   4292210 ns/op	  519510 B/op	   67965 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_10_percent-8    	    2887	    787391 ns/op	  917025 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_20_percent-8    	    3001	    791906 ns/op	  917041 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_30_percent-8    	    2739	    799092 ns/op	  917037 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_40_percent-8    	    3015	    811062 ns/op	  917044 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_50_percent-8    	    2818	    822780 ns/op	  917035 B/op	    6011 allocs/op
Benchmark_Unmarshal_Jsonvalue_ReadOneChain_Blob_60_percent-8    	    2863	    827401 ns/op	  917031 B/op	    6011 allocs/op
PASS
ok  	go.jsonvalue-test/benchmark_test	135.964s
```
