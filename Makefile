.DEFAULT_GOAL := help

# Show this help.
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run: ## run Grpc server and proxy whithout docker
	go run cmd/main.go

up: ## up docker compose
	docker-compose up -d

down: ## down docker compose
	docker-compose down
vendor: ## dowload all dependecies
	go mod vendor
	cd proto && glide install
# general
vendor-proto: ## dowload all dependecies
	cd proto && glide install;

# coverage
WORKDIR = $(PWD)
COVERAGE_REPORT = coverage.txt
COVERAGE_PROFILE = profile.out
COVERAGE_MODE = atomic
MY_VAR := $(shell find $(WORKDIR) -name "*.go" -not -path "*/vendor/*" -not -path "*/gen/*" | grep -o '.*/' |uniq  | xargs)

coverage: ## Coverage test
	@cd $(WORKDIR); \
	go test -v  -race ${MY_VAR} -coverprofile=$(COVERAGE_PROFILE) -covermode=$(COVERAGE_MODE);
compile: ## compile protobuffer
	@cd $(WORKDIR)/proto/src; \
	prototool cache delete;
	@cd $(WORKDIR)/proto/src; \
	prototool generate;
all-docker: ## init all and play witch docker composer
	make vendor-proto && make compile && make vendor && make up ;
all: ## init all and play witch docker composer
	make vendor-proto && make compile && make vendor && make run ;
stop: ## delete all docker composer
	make down;