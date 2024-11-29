ARTIFACTS_DIR = bin

.PHONY: build

build:
	sam build

build-helloWorld:
	mkdir -p $(ARTIFACTS_DIR)/hello-world
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ARTIFACTS_DIR)/hello-world/bootstrap ./hello-world/main.go

.PHONY: invoke
invoke-helloWorld: build-helloWorld
	sam local invoke HelloWorldFunction --template template.yaml
