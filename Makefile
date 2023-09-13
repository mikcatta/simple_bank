sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb    --go_opt=paths=source_relative \
	--go-grpc_out=pb                         --go-grpc_opt=paths=source_relative  \
	--grpc-gateway_out=pb					 --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger 			 --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto 

evans:
	evans --host localhost --port 9090  -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: sqlc test server proto evans redis