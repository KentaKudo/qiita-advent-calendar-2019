SERVICE := qiita-advent-calendar-2019

GIT_HASH := $(shell git rev-parse HEAD)
LINKFLAGS := -X main.gitHash=$(GIT_HASH)

.PHONY: install
install:
	go get -v ./...

$(SERVICE):
	go build -ldflags '$(LINKFLAGS)' .	

.PHONY: clean
clean:
	@rm -f $(SERVICE)
