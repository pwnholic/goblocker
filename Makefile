build:
	@go build -o  bin/goblocker

run: build
	@./bin/goblocker

test: 
	@go test -v ./...

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHOANY: proto
