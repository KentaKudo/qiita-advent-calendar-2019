SERVICE := qiita-advent-calendar-2019

GIT_HASH := $(shell git rev-parse HEAD)
LINKFLAGS := -X main.gitHash=$(GIT_HASH)

.PHONY: install
install:
	go get -v ./...

LINTER_EXE := golangci-lint
LINTER := $(GOPATH)/bin/$(LINTER_EXE)

$(LINTER):
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

LINT_FLAGS :=--enable golint,unconvert,unparam,gofmt

.PHONY: lint
lint: $(LINTER)
	$(LINTER) run $(LINT_FLAGS)

$(SERVICE):
	go build -ldflags '$(LINKFLAGS)' .	

.PHONY: clean
clean:
	@rm -f $(SERVICE)
