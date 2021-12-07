lint: fmt
	golangci-lint run ./... -v -c .golangci.yml

fmt:
	gofumpt -l -w .
	gci -w -local github.com/xenking/binance-api .
