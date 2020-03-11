build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/cli ./cmd/main.go