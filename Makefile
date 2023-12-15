BINARY_NAME=app

.PHONY:
	build \
	dbrun \
	dbstop \
	memorun \
	test \
	gen 

gen: 
	mkdir -p internal
	protoc --go_out=internal --go_opt=paths=source_relative \
		--go-grpc_out=internal --go-grpc_opt=paths=source_relative \
		api/shortener_v1/service.proto

build:
	go build -o $(BINARY_NAME) cmd/api_server/main.go

dbrun: 
	sudo docker compose up -d

dbstop: 
	sudo docker compose down

memorun: build
	./app -config=$(config) -db=false