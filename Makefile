docker-build-run:
	docker build -t vwap-coin .
	docker run -i vwap-coin

local-build-run:
	go build -o vwap-coinbase ./cmd/main.go
	./vwap-coinbase

test:
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	$(MAKE) coverage

coverage:
	go tool cover -func=count.out
