GO=go
GOCOVER=$(GO) tool cover
GOTEST=${GO} test
COVERAGE_DIR=testing
.PHONY: test lint

test:
	$(GOTEST) -v -coverprofile=$(COVERAGE_DIR)/coverage.out contactform
	$(GOCOVER) -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html

coverage:
	open $(COVERAGE_DIR)/coverage.html

lint:
	golangci-lint run --disable unused --disable deadcode