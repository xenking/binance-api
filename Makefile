lint: fmt  ## Run linter for the code
	golangci-lint run ./... -v -c .golangci.yml

fmt:  ## gofmt and goimports all go files
	gci write -s standard -s default -s "prefix(xenking/binance-api)" --skip-generated .
	gofumpt -l -w .

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help