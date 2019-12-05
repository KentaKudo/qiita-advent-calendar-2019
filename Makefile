SERVICE := qiita-advent-calendar-2019

GIT_HASH := $(shell git rev-parse HEAD)
LINKFLAGS := -X main.gitHash=$(GIT_HASH)

PROTO_DIR := ./proto
GENERATED_DIR := ./internal/pb
GENERATED_SERVICE_DIR := $(GENERATED_DIR)/service
GENERATED_EVENT_DIR := $(GENERATED_DIR)/event
GENERATED_ENVELOPE_DIR := $(GENERATED_DIR)/envelope
ENVELOPE_PROTO_MAPPINGS := Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types

.PHONY: protos
protos:
	mkdir -pv $(GENERATED_DIR) $(GENERATED_SERVICE_DIR) $(GENERATED_EVENT_DIR) $(GENERATED_ENVELOPE_DIR)
	protoc \
		-I $(PROTO_DIR) \
		-I $(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=plugins=grpc:$(GENERATED_SERVICE_DIR) \
		service.proto
	protoc \
		-I $(PROTO_DIR) \
		-I $(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=paths=source_relative:$(GENERATED_EVENT_DIR) \
		event.proto
	protoc \
	  	-I $(PROTO_DIR) \
		-I $(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=paths=source_relative,$(ENVELOPE_PROTO_MAPPINGS):$(GENERATED_ENVELOPE_DIR) \
		envelope.proto


mockgen-install:
	GO111MODULE=off go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

mockgen:
	mockgen -source=server.go -package=main -destination server_mock.go -mock_names todoManager=MockTodoManager


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

TEST_FLAGS := -v -cover -timeout 30s

.PHONY: test
test:
	go test $(TEST_FLAGS) ./...

$(SERVICE):
	go build -ldflags '$(LINKFLAGS)' .	

.PHONY: build
build: $(SERVICE)

.PHONY: clean
clean:
	@rm -f $(SERVICE)

.PHONY: all
all: install lint test clean build


DOCKER_REGISTRY=docker.io
DOCKER_REPOSITORY_NAMESPACE=kentakudo
DOCKER_REPOSITORY_IMAGE=$(SERVICE)
DOCKER_REPOSITORY=$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY_NAMESPACE)/$(DOCKER_REPOSITORY_IMAGE)
DOCKER_IMAGE_TAG=$(GIT_HASH)

.PHONY: docker-image
docker-image:
	docker build -t $(DOCKER_REPOSITORY):$(DOCKER_IMAGE_TAG) . \
	  --build-arg SERVICE=$(SERVICE)

.PHONY: docker-auth
docker-auth:
	@docker login -u $(DOCKER_ID) -p $(DOCKER_PASSWORD) $(DOCKER_REGISTRY)

.PHONY: docker-build
docker-build: docker-image docker-auth
	docker tag $(DOCKER_REPOSITORY):$(DOCKER_IMAGE_TAG) $(DOCKER_REPOSITORY):latest
	docker push $(DOCKER_REPOSITORY)


K8S_NAMESPACE=qiita
K8S_DEPLOYMENT_NAME=$(SERVICE)
K8S_CONTAINER_NAME=$(SERVICE)
K8S_URL=https://<dev env>/apis/apps/v1/namespaces/$(K8S_NAMESPACE)/deployments/$(K8S_DEPLOYMENT_NAME)
K8S_PAYLOAD={"spec":{"template":{"spec":{"containers":[{"name":"$(K8S_CONTAINER_NAME)","image":"$(DOCKER_REPOSITORY):$(DOCKER_IMAGE_TAG)"}]}}}}

.PHONY: kubernetes-push
kubernetes-push:
	test "$(shell curl -o /dev/null -w '%{http_code}' -s -X PATCH -k -d '$(K8S_PAYLOAD)' -H 'Content-Type: application/strategic-merge-patch+json' -H 'Authorization: Bearer $(K8S_AUTH_TOKEN)' '$(K8S_URL)')" -eq "200"

