DB_SOURCE="postgresql://postgres:password@localhost/db_latihan_gin_prod?sslmode=disable&search_path=db_latihan_gin_prod"
MIGRATIONS_DIR=./migrations

# make migrate-init name="tbl_users"
migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql ${name}

# make migrate-up
migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

# make migrate-down
migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

# make migrate-fix
migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0
