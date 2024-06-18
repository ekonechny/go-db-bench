# Golang SQL-frameworks benchmarks

Comparing the performanc
e of SQL-frameworks

### Select queries
```
======= limit:1 =========
Benchmark/database/sql/libpq_limit:1-8         	    4322	    862162 ns/op	    2528 B/op	      62 allocs/op
Benchmark/sqlc/pgx_limit:1-8                   	    8272	    437183 ns/op	    1472 B/op	      31 allocs/op
Benchmark/sqlc/libpq_limit:1-8                 	    4075	    859805 ns/op	    2560 B/op	      64 allocs/op
Benchmark/xo/libpq_limit:1-8                   	    4278	    867682 ns/op	    2408 B/op	      64 allocs/op
Benchmark/sqlx/libpq_limit:1-8                 	    4095	    867736 ns/op	    3040 B/op	      68 allocs/op
Benchmark/gorm/pgx_limit:1-8                   	    7816	    432242 ns/op	    9037 B/op	     118 allocs/op
Benchmark/bun_limit:1-8                        	    8190	    444572 ns/op	    7200 B/op	      64 allocs/op

======= limit:10 =========
Benchmark/database/sql/libpq_limit:10-8        	    3894	    905953 ns/op	   17808 B/op	     404 allocs/op
Benchmark/sqlc/pgx_limit:10-8                  	    7436	    466017 ns/op	   12824 B/op	     181 allocs/op
Benchmark/sqlc/libpq_limit:10-8                	    3919	    902169 ns/op	   17840 B/op	     406 allocs/op
Benchmark/xo/libpq_limit:10-8                  	    3921	    910453 ns/op	   13000 B/op	     406 allocs/op
Benchmark/sqlx/libpq_limit:10-8                	    3914	    914557 ns/op	   18536 B/op	     419 allocs/op
Benchmark/gorm/pgx_limit:10-8                  	    7018	    498939 ns/op	   22546 B/op	     591 allocs/op
Benchmark/bun_limit:10-8                       	    7134	    495858 ns/op	   20776 B/op	     387 allocs/op

======= limit:100 =========
Benchmark/database/sql/libpq_limit:100-8       	    3067	   1160005 ns/op	  157889 B/op	    3777 allocs/op
Benchmark/sqlc/pgx_limit:100-8                 	    5011	    676867 ns/op	  111546 B/op	    1645 allocs/op
Benchmark/sqlc/libpq_limit:100-8               	    3010	   1171812 ns/op	  157921 B/op	    3779 allocs/op
Benchmark/xo/libpq_limit:100-8                 	    3054	   1152297 ns/op	  116986 B/op	    3779 allocs/op
Benchmark/sqlx/libpq_limit:100-8               	    3014	   1174343 ns/op	  160777 B/op	    3882 allocs/op
Benchmark/gorm/pgx_limit:100-8                 	    4286	    828978 ns/op	  198775 B/op	    5316 allocs/op
Benchmark/bun_limit:100-8                      	    4569	    763487 ns/op	  144553 B/op	    3613 allocs/op

======= limit:1000 =========
Benchmark/database/sql/libpq_limit:1000-8      	    1071	   3319578 ns/op	 1484020 B/op	   37816 allocs/op
Benchmark/sqlc/pgx_limit:1000-8                	    1131	   3076066 ns/op	 1251873 B/op	   16373 allocs/op
Benchmark/sqlc/libpq_limit:1000-8              	    1075	   3846425 ns/op	 1484163 B/op	   37822 allocs/op
Benchmark/xo/libpq_limit:1000-8                	     988	   3494709 ns/op	 1155459 B/op	   37821 allocs/op
Benchmark/sqlx/libpq_limit:1000-8              	     897	   3819124 ns/op	 1508557 B/op	   38822 allocs/op
Benchmark/gorm/pgx_limit:1000-8                	     934	   4196463 ns/op	 1843922 B/op	   52876 allocs/op
Benchmark/bun_limit:1000-8                     	    1086	   3243560 ns/op	 1302949 B/op	   35899 allocs/op

======= limit:10000 =========
Benchmark/database/sql/libpq_limit:10000-8     	     120	  31943042 ns/op	17847692 B/op	  377807 allocs/op
Benchmark/sqlc/pgx_limit:10000-8               	     126	  24221654 ns/op	15268065 B/op	  163480 allocs/op
Benchmark/sqlc/libpq_limit:10000-8             	     128	  33427027 ns/op	17847315 B/op	  377796 allocs/op
Benchmark/xo/libpq_limit:10000-8               	     147	  30945900 ns/op	11675746 B/op	  377800 allocs/op
Benchmark/sqlx/libpq_limit:10000-8             	     100	  33812548 ns/op	18087967 B/op	  387804 allocs/op
Benchmark/gorm/pgx_limit:10000-8               	      98	  34374956 ns/op	21365301 B/op	  527974 allocs/op
Benchmark/bun_limit:10000-8                    	     100	  34775032 ns/op	15993365 B/op	  358622 allocs/op

======= limit:100000 =========
Benchmark/database/sql/libpq_limit:100000-8    	      64	 234582340 ns/op	190785841 B/op	 3777536 allocs/op
Benchmark/sqlc/pgx_limit:100000-8              	      64	 186509064 ns/op	164712218 B/op	 1634507 allocs/op
Benchmark/sqlc/libpq_limit:100000-8            	      55	 202134434 ns/op	190785682 B/op	 3777535 allocs/op
Benchmark/xo/libpq_limit:100000-8              	      37	 317997885 ns/op	118135792 B/op	 3777536 allocs/op
Benchmark/sqlx/libpq_limit:100000-8            	      51	 275511416 ns/op	193186085 B/op	 3877540 allocs/op
Benchmark/gorm/pgx_limit:100000-8              	      38	 317811727 ns/op	225644409 B/op	 5277849 allocs/op
Benchmark/bun_limit:100000-8                   	      15	 258468775 ns/op	172202384 B/op	 3585734 allocs/op
```

### Insert queries
```

======= insert:1 =========
Benchmark/database/sql/libpq_batchSize:1-8     	    2160	   1604343 ns/op	   67912 B/op	      49 allocs/op
Benchmark/sqlc/pgx_batchSize:1-8               	    5595	    620827 ns/op	    2511 B/op	      56 allocs/op
Benchmark/sqlc/libpq_batchSize:1-8             	    3938	    935750 ns/op	    2160 B/op	      37 allocs/op
Benchmark/sqlx/libpq_batchSize:1-8             	    3831	    922366 ns/op	    2688 B/op	      53 allocs/op
Benchmark/gorm/pgx_batchSize:1-8               	    2997	   1211255 ns/op	   10883 B/op	     199 allocs/op
Benchmark/bun_batchSize:1-8                    	    6141	    560510 ns/op	    6360 B/op	      16 allocs/op

======= insert:10 =========
Benchmark/database/sql/libpq_batchSize:10-8    	    2016	   1784021 ns/op	   77757 B/op	     247 allocs/op
Benchmark/sqlc/pgx_batchSize:10-8              	    4705	    763845 ns/op	   14074 B/op	     283 allocs/op
Benchmark/sqlc/libpq_batchSize:10-8            	    3196	   1163995 ns/op	   22312 B/op	     220 allocs/op
Benchmark/sqlx/libpq_batchSize:10-8            	    2985	   1203603 ns/op	   25573 B/op	     261 allocs/op
Benchmark/gorm/pgx_batchSize:10-8              	    2481	   1460720 ns/op	   79619 B/op	    1399 allocs/op
Benchmark/bun_batchSize:10-8                   	    4038	    981959 ns/op	    9096 B/op	      35 allocs/op

======= insert:100 =========
Benchmark/database/sql/libpq_batchSize:100-8   	     974	   3670629 ns/op	  175477 B/op	    2213 allocs/op
Benchmark/sqlc/pgx_batchSize:100-8             	    2229	   2285321 ns/op	  165403 B/op	    2503 allocs/op
Benchmark/sqlc/libpq_batchSize:100-8           	    1458	   2477757 ns/op	  251166 B/op	    2012 allocs/op
Benchmark/sqlx/libpq_batchSize:100-8           	    1389	   2556781 ns/op	  284379 B/op	    2240 allocs/op
Benchmark/gorm/pgx_batchSize:100-8             	    1244	   3223911 ns/op	  767752 B/op	   13759 allocs/op
Benchmark/bun_batchSize:100-8                  	    1666	   1936719 ns/op	  103784 B/op	     221 allocs/op

======= insert:1000 =========
Benchmark/database/sql/libpq_batchSize:1000-8  	     384	  10132914 ns/op	 1138941 B/op	   21864 allocs/op
Benchmark/sqlc/pgx_batchSize:1000-8            	     462	   7963290 ns/op	  793136 B/op	   24143 allocs/op
Benchmark/sqlc/libpq_batchSize:1000-8          	     272	  11734368 ns/op	 2524820 B/op	   19832 allocs/op
Benchmark/sqlx/libpq_batchSize:1000-8          	     268	  13832112 ns/op	 2922497 B/op	   21896 allocs/op
Benchmark/gorm/pgx_batchSize:1000-8            	     211	  18206617 ns/op	 7964807 B/op	  134505 allocs/op
Benchmark/bun_batchSize:1000-8                 	     277	  11872081 ns/op	 1154616 B/op	    2030 allocs/op
```