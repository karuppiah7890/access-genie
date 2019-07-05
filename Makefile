build: vendor
	CGO_ENABLED=0 go build -o out/access-genie -v -mod=vendor bot/main.go
	CGO_ENABLED=0 go build -o out/access-genie-service -v -mod=vendor service/main.go

setup:
	cp .env.sample .env

test:
	CGO_ENABLED=0 go test ./... -cover

vendor:
	go mod vendor -v