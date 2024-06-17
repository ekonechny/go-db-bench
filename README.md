# Golang SQL-frameworks benchmarks

Comparing the performanc
e of SQL-frameworks

```
======= limit:1 =========
goos: darwin
goarch: arm64
pkg: go-db-bench
Benchmark/database/sql/libpq_limit:1-8              1220            939551 ns/op            1328 B/op         33 allocs/op
Benchmark/sqlx/libpq_limit:1-8                      1256            932980 ns/op            1680 B/op         39 allocs/op
Benchmark/gorm/pgx_limit:1-8                        2746            508742 ns/op            5354 B/op         71 allocs/op
Benchmark/sqlc/pgx_limit:1-8                        2886            458637 ns/op             768 B/op         14 allocs/op
Benchmark/sqlc/libpq_limit:1-8                      1452            937499 ns/op            1360 B/op         35 allocs/op
Benchmark/xo/libpq_limit:1-8                        1314            917164 ns/op            1304 B/op         35 allocs/op

======= limit:10 =========
Benchmark/database/sql/libpq_limit:10-8             1263            961092 ns/op            5792 B/op        134 allocs/op
Benchmark/sqlx/libpq_limit:10-8                     1314            959342 ns/op            6360 B/op        149 allocs/op
Benchmark/gorm/pgx_limit:10-8                       2229            556202 ns/op            8151 B/op        222 allocs/op
Benchmark/sqlc/pgx_limit:10-8                       2590            486424 ns/op            4688 B/op         45 allocs/op
Benchmark/sqlc/libpq_limit:10-8                     1294            949443 ns/op            5824 B/op        136 allocs/op
Benchmark/xo/libpq_limit:10-8                       1369            944730 ns/op            3960 B/op        136 allocs/op

======= limit:100 =========
Benchmark/database/sql/libpq_limit:100-8            1022           1124062 ns/op           46784 B/op       1098 allocs/op
Benchmark/sqlx/libpq_limit:100-8                    1207           1170677 ns/op           49512 B/op       1203 allocs/op
Benchmark/gorm/pgx_limit:100-8                      1428            931035 ns/op           55099 B/op       1726 allocs/op
Benchmark/sqlc/pgx_limit:100-8                      2278            573437 ns/op           39536 B/op        318 allocs/op
Benchmark/sqlc/libpq_limit:100-8                    1104           1121037 ns/op           46816 B/op       1100 allocs/op
Benchmark/xo/libpq_limit:100-8                      1071           1095417 ns/op           29976 B/op       1100 allocs/op

======= limit:1000 =========
Benchmark/database/sql/libpq_limit:1000-8            504           2672430 ns/op          414555 B/op      10674 allocs/op
Benchmark/sqlx/libpq_limit:1000-8                    326           3968984 ns/op          438884 B/op      11679 allocs/op
Benchmark/gorm/pgx_limit:1000-8                      206           6428040 ns/op          463011 B/op      16703 allocs/op
Benchmark/sqlc/pgx_limit:1000-8                      938           1294710 ns/op          351516 B/op       3023 allocs/op
Benchmark/sqlc/libpq_limit:1000-8                    489           2428024 ns/op          414587 B/op      10676 allocs/op
Benchmark/xo/libpq_limit:1000-8                      530           2584087 ns/op          286145 B/op      10677 allocs/op

======= limit:10000 =========
Benchmark/database/sql/libpq_limit:10000-8           165           7140782 ns/op         5542918 B/op     106507 allocs/op
Benchmark/sqlx/libpq_limit:10000-8                   158           7545367 ns/op         5783283 B/op     116512 allocs/op
Benchmark/gorm/pgx_limit:10000-8                     100          10506602 ns/op         5983186 B/op     166572 allocs/op
Benchmark/sqlc/pgx_limit:10000-8                     169           7065403 ns/op         5161680 B/op      30031 allocs/op
Benchmark/sqlc/libpq_limit:10000-8                   168           6891126 ns/op         5542955 B/op     106509 allocs/op
Benchmark/xo/libpq_limit:10000-8                     202           6116851 ns/op         2987619 B/op     106509 allocs/op

======= limit:100000 =========
Benchmark/database/sql/libpq_limit:100000-8           19          59439544 ns/op        61273492 B/op    1063917 allocs/op
Benchmark/sqlx/libpq_limit:100000-8                   16          72810612 ns/op        63673714 B/op    1163921 allocs/op
Benchmark/gorm/pgx_limit:100000-8                     10         105651417 ns/op        65553063 B/op    1664112 allocs/op
Benchmark/sqlc/pgx_limit:100000-8                     16          66882456 ns/op        58157408 B/op     300041 allocs/op
Benchmark/sqlc/libpq_limit:100000-8                   19          61324675 ns/op        61273463 B/op    1063918 allocs/op
Benchmark/xo/libpq_limit:100000-8                     19          55774121 ns/op        31258630 B/op    1063919 allocs/op
```