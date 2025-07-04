dev:
	export CONFIG_PATH=config/local.yml && air

swagger:
	swag fmt && swag init -g cmd/kushfinds/main.go

gen:
	go generate ./...

cover:
	go test -count=1 -race -coverprofile=coverage.out ./internal/...
	go tool cover -html=coverage.out
	rm coverage.out