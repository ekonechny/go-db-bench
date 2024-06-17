# Golang SQL-frameworks benchmarks

Comparing the performanc
e of SQL-frameworks

```
======= limit:1 =========
Benchmark/database/sql/libpq_limit:1-8              1179            880178 ns/op            1328 B/op         33 allocs/op
Benchmark/sqlx/libpq_limit:1-8                      1404            902760 ns/op            1680 B/op         39 allocs/op
Benchmark/gorm/pgx_limit:1-8                        2568            448074 ns/op            5354 B/op         71 allocs/op
Benchmark/sqlc/pgx_limit:1-8                        2997            450209 ns/op             768 B/op         14 allocs/op
Benchmark/sqlc/libpq_limit:1-8                      1429            877707 ns/op            1360 B/op         35 allocs/op
Benchmark/xo/libpq_limit:1-8                        1358            886210 ns/op            1304 B/op         35 allocs/op

======= limit:10 =========
Benchmark/database/sql/libpq_limit:10-8             1438            920418 ns/op            5792 B/op        134 allocs/op
Benchmark/sqlx/libpq_limit:10-8                     1334            958959 ns/op            6360 B/op        149 allocs/op
Benchmark/gorm/pgx_limit:10-8                       2122            554075 ns/op            8154 B/op        222 allocs/op
Benchmark/sqlc/pgx_limit:10-8                       2698            488331 ns/op            4688 B/op         45 allocs/op
Benchmark/sqlc/libpq_limit:10-8                     1312            957336 ns/op            5824 B/op        136 allocs/op
Benchmark/xo/libpq_limit:10-8                       1281            942170 ns/op            3960 B/op        136 allocs/op

======= limit:100 =========
Benchmark/database/sql/libpq_limit:100-8            1053           1163303 ns/op           46784 B/op       1098 allocs/op
Benchmark/sqlx/libpq_limit:100-8                    1087           1157892 ns/op           49512 B/op       1203 allocs/op
Benchmark/gorm/pgx_limit:100-8                      1444           1082700 ns/op           55104 B/op       1726 allocs/op
Benchmark/sqlc/pgx_limit:100-8                      2035            585650 ns/op           39536 B/op        318 allocs/op
Benchmark/sqlc/libpq_limit:100-8                     973           1120088 ns/op           46816 B/op       1100 allocs/op
Benchmark/xo/libpq_limit:100-8                      1148           1140068 ns/op           29976 B/op       1100 allocs/op

======= limit:1000 =========
Benchmark/database/sql/libpq_limit:1000-8            558           2498577 ns/op          414554 B/op      10674 allocs/op
Benchmark/sqlx/libpq_limit:1000-8                    391           3424246 ns/op          438883 B/op      11679 allocs/op
Benchmark/gorm/pgx_limit:1000-8                      414           2995523 ns/op          463039 B/op      16704 allocs/op
Benchmark/sqlc/pgx_limit:1000-8                     1125           1223704 ns/op          351514 B/op       3023 allocs/op
Benchmark/sqlc/libpq_limit:1000-8                    567           2212690 ns/op          414587 B/op      10676 allocs/op
Benchmark/xo/libpq_limit:1000-8                      525           2367505 ns/op          286146 B/op      10677 allocs/op

======= limit:10000 =========
Benchmark/database/sql/libpq_limit:10000-8           177           7024442 ns/op         5542931 B/op     106507 allocs/op
Benchmark/sqlx/libpq_limit:10000-8                   157           8017937 ns/op         5783254 B/op     116512 allocs/op
Benchmark/gorm/pgx_limit:10000-8                     100          10508128 ns/op         5983646 B/op     166576 allocs/op
Benchmark/sqlc/pgx_limit:10000-8                     196           6019278 ns/op         5161693 B/op      30031 allocs/op
Benchmark/sqlc/libpq_limit:10000-8                   165           7130181 ns/op         5542946 B/op     106509 allocs/op
Benchmark/xo/libpq_limit:10000-8                     160           7069956 ns/op         2987600 B/op     106509 allocs/op

======= limit:100000 =========
Benchmark/database/sql/libpq_limit:100000-8           18          65571546 ns/op        61273444 B/op    1063916 allocs/op
Benchmark/sqlx/libpq_limit:100000-8                   16          72419773 ns/op        63673719 B/op    1163921 allocs/op
Benchmark/gorm/pgx_limit:100000-8                     10         105282096 ns/op        65552640 B/op    1664115 allocs/op
Benchmark/sqlc/pgx_limit:100000-8                     27          43653758 ns/op        58157459 B/op     300041 allocs/op
Benchmark/sqlc/libpq_limit:100000-8                   18          66917148 ns/op        61273458 B/op    1063918 allocs/op
Benchmark/xo/libpq_limit:100000-8                     18          65066241 ns/op        31258645 B/op    1063919 allocs/op

```