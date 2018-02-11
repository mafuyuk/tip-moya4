.DEFAULT_GOAL:=all-check
HAVE_GOLINT:=$(shell which golint)
HAVE_DEP:=$(shell which dep)
HAVE_ENUMER:=$(shell which enumer)
HAVE_PECO:=$(shell which peco)
HAVE_MIGRATE:=$(shell which migrate)
HAVE_CIRCLECI:=$(shell which circleci)

.PHONY: setup
setup: dep enumer
	@echo "go setup"
	@dep ensure
	@go generate ./...

## Check
.PHONY: all-check lint vet test
all-check: lint vet test
	@echo "all check"

lint: setup golint
	@echo "go lint"
	@golint $(shell go list ./...|grep -v vendor)

vet: setup
	@echo "go vet"
	@go vet ./...

test: setup
	@echo "go test"
	@go test -v ./...


## Docker
CONTAINER_PREFIX:=tip-moya4

.PHONY: dstart dstop dstatus dlogin dclean dlog
dstart: setup
	@echo "docker start"
	@docker-compose up -d

dstop:
	@echo "docker stop"
	@docker-compose stop

dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/bash

dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force
	@rm -rf ./minio-data

dlog: peco
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)

## CI
.PHONY: cibuild cilint
cibuild: circleci
	@echo "circle ci local build"
	@circleci build

cilint: circleci
	@echo "circle ci local lint check"
	@circleci config validate -c ./.circleci/config.yml

## Go package
.PHONY: dep enumer golint peco migrate
dep:
ifndef HAVE_DEP
	@echo "Installing dep"
	@go get -u github.com/golang/dep/cmd/dep
endif

enumer:
ifndef HAVE_ENUMER
	@echo "Installing enumer"
	@go get -u github.com/alvaroloes/enumer
endif

golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

peco:
ifndef HAVE_PECO
	@echo "Installing peco"
	@go get -u github.com/peco/peco/cmd/peco
endif

migrate:
ifndef HAVE_MIGRATE
	@echo "Installing migrate"
	@go get -u -d github.com/mattes/migrate/cli github.com/go-sql-driver/mysql
	@go build -tags 'mysql' -o ${GOPATH}/bin/migrate github.com/mattes/migrate/cli
endif

circleci:
ifndef HAVE_CIRCLECI
	@echo "Installing circleci"
	@curl -o /usr/local/bin/circleci https://circle-downloads.s3.amazonaws.com/releases/build_agent_wrapper/circleci
	@chmod +x /usr/local/bin/circleci
endif
