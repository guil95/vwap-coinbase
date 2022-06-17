docker-build-run:
	docker build -t vwap-coin .
	docker run -i vwap-coin

local-build-run:
	go mod vendor
	go build -o vwap-coinbase ./cmd/main.go
	./vwap-coinbase

test:
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	$(MAKE) coverage

coverage:
	go tool cover -func=count.out

lint: ## Run linter
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.44.0 golangci-lint run -v
