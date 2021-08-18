BUILD_SETTINGS = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
IMAGE = book-server


.PHONY: all build coverage clean fmt release test vet

all: build coverage clean

server: build 
	@./book-server
	
build: fmt vet test
	$(BUILD_SETTINGS) go build -trimpath -o "$(IMAGE)" ./main.go

fmt:
	@go fmt ./...

test:
	@go test -race $(shell go list ./... | grep -v e2e) --coverprofile=coverage.out

coverage: test
	@go tool cover -func=coverage.out

vet:
	@go vet ./...

clean:
	@rm -rf ./$(IMAGE)
	@go clean -cache
	@rm -rf *.out
