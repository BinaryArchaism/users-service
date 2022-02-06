gen:
	protoc --go-grpc_out=../models/ ../users.proto

mig_create:
	migrate create -ext sql -dir db-migrations -seq ${NAME}

mig_up:
	migrate -database ${URL} -path db-migrations/postgres up

mig_down:
	migrate -database ${URL} -path db-migrations/postgres down

clh:
	migrate -database clickhouse://localhost:8123?username=admin&password=test&database=default -path db-migrations/clickhouse up
	migrate -database clickhouse://localhost:8123 -path db-migrations/clickhouse up
