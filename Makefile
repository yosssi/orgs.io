run:
	go run cmd/orgs.io/server.go -c orgs.io.yaml
test:
	@go test -cover ./app/controllers
	@go test -cover ./app/models
	@go test -cover ./app/router
	@go test -cover ./cmd/orgs.io
