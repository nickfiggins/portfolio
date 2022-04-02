GO=go
GOCOVER=$(GO) tool cover
GOTEST=${GO} test
.PHONY: test

test:
	$(GOTEST) -v -coverprofile=coverage.out contactform
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out