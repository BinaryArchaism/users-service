gen:
	protoc --go-grpc_out=../models/ ../users.proto

mig_create:
	migrate create -ext sql -dir db-migrations -seq ${NAME}

mig_up:
	migrate -database ${URL} -path db-migrations up

mig_down:
	migrate -database ${URL} -path db-migrations down