r:
	go run cmd/orgs.io/server.go -c orgs.io.yml
t:
	@go test -cover -race ./app/controllers
	@go test -cover -race ./app/models
	@go test -cover -race ./app/router
	@go test -cover -race ./cmd/orgs.io
