sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb    --go_opt=paths=source_relative \
	--go-grpc_out=pb                         --go-grpc_opt=paths=source_relative  \
	--grpc-gateway_out=pb					 --grpc-gateway_opt=paths=source_relative \
	proto/*.proto 

evans:
	evans --host localhost --port 9090  -r repl

.PHONY: sqlc test server proto evans