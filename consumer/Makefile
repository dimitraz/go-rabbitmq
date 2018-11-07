DOCKER_ORG=dimitraz
IMAGE_NAME=rabbitmq-consumer
IMAGE_TAG=1.1
BINARY_NAME=consumer
GOOS=linux

.phony: docker_release
docker_release: docker_build docker_push

.phony: docker_push
docker_push:
	docker push $(DOCKER_ORG)/$(IMAGE_NAME):$(IMAGE_TAG)

.phony: docker_build 
docker_build: build_binary
	docker build -t $(DOCKER_ORG)/$(IMAGE_NAME):$(IMAGE_TAG) .

.phony: build_binary
build_binary:
	GOOS=$(GOOS) go build -o $(BINARY_NAME) main.go