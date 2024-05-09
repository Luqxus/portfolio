build:
	@go build -o ./bin/profile

run: build
	@./bin/profile

test:
	@go test ./...