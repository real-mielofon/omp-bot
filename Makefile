SERVICE_NAME=omp-bot
SERVICE_PATH=real-mielofon/omp-bot

SERVICE1_NAME=rtg-service-api
SERVICE1_PATH=ozonmp/rtg-service-api
SERVICE2_NAME=rtg-service-facade
SERVICE2_PATH=ozonmp/rtg-service-facade

PGV_VERSION:="v0.6.1"
BUF_VERSION:="v0.56.0"

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin
GOEXE=GO_BIN
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

ifeq ("NT", "$(findstring NT,$(OS_NAME))")
OS_NAME=Windows
endif


.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o bot cmd/bot/main.go



.PHONY: generate
generate: .generate-go .generate-finalize-go

.PHONY: generate
generate-go: .generate-go .generate-finalize-go

.generate-go:
	$(BUF_EXE) generate

.generate-finalize-go:
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE1_PATH)/pkg/$(SERVICE1_NAME)/* pkg/$(SERVICE1_NAME)
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE2_PATH)/pkg/$(SERVICE2_NAME)/* pkg/$(SERVICE2_NAME)
	rm -rf pkg/$(SERVICE_NAME)/github.com/
	cd pkg/$(SERVICE1_NAME) && ls go.mod || (go mod init github.com/$(SERVICE1_PATH)/pkg/$(SERVICE1_NAME) && go mod tidy)
	cd pkg/$(SERVICE2_NAME) && ls go.mod || (go mod init github.com/$(SERVICE2_PATH)/pkg/$(SERVICE2_NAME) && go mod tidy)
