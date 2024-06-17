DATABASE_URL := host=127.0.0.1 port=5432 database=postgres user=postgres sslmode=disable
export DATABASE_URL

gen:
	@sqlc --file sqlc/libpq/sqlc.json generate
	@sqlc --file sqlc/pgx/sqlc.json generate
	@xo schema postgres://postgres:postgres@127.0.0.1/postgres?sslmode=disable -o xo
	@xo query postgres://postgres:postgres@127.0.0.1/postgres?sslmode=disable -o xo -M -B -2 -T Products -Q " \
     SELECT * \
     FROM products \
     LIMIT %%limit int%% \
     "


docker-compose:
	@docker-compose -p db-bench up -d

docker-compose-down:
	@docker-compose -p db-bench down

load-dump:
	@psql postgres -U postgres -h localhos < dumpfile

run:
	go test -bench=.  -benchmem ./...