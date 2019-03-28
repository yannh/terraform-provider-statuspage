#!/usr/bin/make

TFPLAN ?= plan.tfplan

all: plan

vet:
	go vet ./...

test: vet
	go test ./...

build: test
	go install

init: build
	terraform init -plugin-dir $(GOPATH)bin

plan: init
	terraform plan -out ${TFPLAN}

apply:
	terraform apply ${TFPLAN}