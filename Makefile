#============ Vendor ====================
vendor:
	go mod vendor -v

#============ Testing ====================
test: 
	go fmt ./...
	go test -shuffle=on ./...

#========== Run Application ==============
server: vendor api
api: 
	go run ./cmd/app/main.go

cli:
	go build ./cmd/cli
	go install ./cmd/cli

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0
