build:
	CGO_ENABLED=0 go build -o out/access-genie -v -mod=vendor bot/main.go

setup:
	cp .env.sample .env

test:
	CGO_ENABLED=0 go test ./... -cover