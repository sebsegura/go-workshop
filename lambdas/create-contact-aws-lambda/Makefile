BIN= $(CURDIR)/bin
Q= $(if $(filter 1,$V),,@)
M= $(shell printf "\033[34;1m▶\033[0m")
FUNCTION=create-contact-aws-lambda

all: build

deps: ;(info $(M) installing dependencies...) @
	@go mod tidy

build: ; $(info $(M) building executable...) @
	@env GOOS=linux go build -ldflags="-s -w" -o bin/$(FUNCTION) cmd/main.go

lint:
	@go vet ./...

format:
	@go fmt ./...

clean:
	@rm -rf $(BIN)

zip: build
	@zip -j $(FUNCTION).zip bin/$(FUNCTION)

.PHONY: test
test:
	$(HOME)/go/bin/ginkgo -r --progress --failFast  --randomizeAllSpecs --randomizeSuites --failOnPending --trace --reportFile=./junit.xml -coverpkg=./... -coverprofile=coverage.out -outputdir=./test

.PHONY: report
report:test
   	@go tool cover -html=./test/coverage.out -o ./test/coverage.html


.PHONY: build
