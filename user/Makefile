DB="postgres://postgres:2005@localhost:5432/udevs?sslmode=disable"

run:
	go run cmd/main.go

table:
	migrate create -dir migrations -ext sql db
table_up:
	migrate -path migrations -database ${DB} up
table_down:
	migrate -path migrations -database ${DB} down
table_force:
	migrate -path migrations -database ${DB} force 20250124150744.
gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./internal/protos/user.proto