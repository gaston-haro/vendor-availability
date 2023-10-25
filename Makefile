BINARY=service

# Actions to execute locally
build:
	go build -o ${BINARY} cmd/web/*.go

clean:
	go clean -testcache ./...

test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

test_local:
	go test -cover ./...

test_local_coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
	
web:
	@clear
	go build -o ${BINARY} cmd/web/*.go
	./${BINARY} -E dev

# Optional actions for projects that may need them
#test_integration:
#	go clean -testcache && go test -v -tags=integration ./test/integration