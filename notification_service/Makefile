DB="postgres://postgres:2005@localhost:5432/udevs1?sslmode=disable"

run:
	go run cmd/main.go
gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./internal/protos/notification.proto