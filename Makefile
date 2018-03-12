GOCMD=go
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mattertime

all: build

build:
	$(GOINSTALL)

test:
	$(GOTEST) -v .

run:
	$(GOINSTALL)
	./$(BINARY_NAME)

# deps:
	#$(GOGET) github.com/golang/dep/cmd/dep
	#dep ensure -update