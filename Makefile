# use linter for formatted code
lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run

# use command for build app
.PHONY: build
build:
	go mod vendor
	go build -ldflags "-w -s" -o misha cmd/misha/main.go
