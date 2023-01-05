# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Project name and binary name
PROJECTNAME=myapp
BINARYNAME=myapp

# Directories
SRCDIR=src
BINDIR=bin

# Main build target
all: test build

# Install dependencies
deps:
	$(GOGET) -u github.com/securego/gosec/cmd/gosec

# Run gofmt on all source files
fmt:
	$(GOCMD) fmt $(SRCDIR)/*.go

# Run gosec on all source files
security: deps
	gosec $(SRCDIR)/*.go

# Run tests
test:
	$(GOTEST) -v $(SRCDIR)/...

# Build the binary
build:
	$(GOBUILD) -o $(BINDIR)/$(BINARYNAME) $(SRCDIR)/main.go

# Clean up
clean:
	$(GOCLEAN)
	rm -rf $(BINDIR)
