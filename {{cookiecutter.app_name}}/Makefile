.PHONY: build build-alpine clean test help default

BIN_NAME={{cookiecutter.app_name}}

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_COMMIT_CNT := $(shell git rev-list --all --count)
GIT_TAG=$(shell git describe --tags --abbrev=0)
GIT_HASH=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

TAG=$(GIT_TAG).$(GIT_COMMIT_CNT).$(GIT_HASH)

default: test

help:
	@echo 'Management commands for {{cookiecutter.app_name}}:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	go build -o bin/${BIN_NAME}

docker-save:
	docker save wangjiaxin/$(BIN_NAME):$(TAG) > $(BIN_NAME).$(TAG).tar && docker save wangjiaxin/$(BIN_NAME):latest > $(BIN_NAME).latest.tar

docker-build:
	docker build . -t wangjiaxin/$(BIN_NAME):$(TAG) && docker tag wangjiaxin/$(BIN_NAME):$(TAG) wangjiaxin/$(BIN_NAME):latest && \
	docker tag wangjiaxin/$(BIN_NAME):$(TAG) registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):$(TAG) && \
	docker tag wangjiaxin/$(BIN_NAME):$(TAG) registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):$(GIT_TAG) && \
	docker tag wangjiaxin/$(BIN_NAME):$(TAG) registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):latest

docker-push:
	docker push wangjiaxin/$(BIN_NAME):$(TAG) && docker push wangjiaxin/$(BIN_NAME):latest

docker-push-aliyun:
	docker push registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):$(TAG); \
	docker push registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):$(GIT_TAG); \
	docker push registry.cn-beijing.aliyuncs.com/showmethemoney/$(BIN_NAME):latest

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

