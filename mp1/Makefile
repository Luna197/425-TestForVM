# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mp1

.PHONY: build test clean run
.default: test

#build: 
#		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GORUN) src/test.go src/utils.go
clean: 
		rm -f $(BINARY_NAME)