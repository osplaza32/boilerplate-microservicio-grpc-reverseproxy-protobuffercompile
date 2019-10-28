.DEFAULT_GOAL := help

# Show this help.
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

prepare: ## task to prepare whatever
	@echo "prepare task triggered"

install: ## task to install whatever
	@echo "install task triggered"
run: ## run Grpc server and proxy whithout docker
	go run cmd/main.go

up: ## up docker compose
	docker-compose up -d

down: ## down docker compose
	docker-compose down
vendor: ## dowload all dependecies
	go mod vendor
	cd proto && glide install
.PHONY: help  prepare install vendor run up down
