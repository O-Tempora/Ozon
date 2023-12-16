BINARY_NAME=app
IMAGE_NAME=api
PORT=6969
APP_PORT=6969

.PHONY:
	build \
	dbrun \
	dbstop \
	memorun \
	memostop \
	test \
	gen \
	run

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

memorun:
	sudo docker build -t $(IMAGE_NAME) .
	sudo docker run -d -p $(PORT):$(APP_PORT) --name $(IMAGE_NAME) -it $(IMAGE_NAME) ./$(BINARY_NAME) -config=config/docker.yaml -db=false

memostop:
	sudo docker stop $(IMAGE_NAME) && sudo docker rm $(IMAGE_NAME)

run: build
	./$(BINARY_NAME) -db=false

test:
	go test -count=1 ./...