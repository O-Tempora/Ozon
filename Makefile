# host=localhost
# port=6969
# dbport=6970
# dbname=ozon
# dbuser=postgres
# dbpass=postgres

.PHONY:
	build \
	dbrun \
	memorun \
	test \
	gen 

gen:
	mkdir -p internal
	protoc --go_out=internal --go_opt=paths=source_relative \
		--go-grpc_out=internal --go-grpc_opt=paths=source_relative \
		api/shortener_v1/service.proto
