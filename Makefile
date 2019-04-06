#!/usr/bin/make

TFPLAN ?= plan.tfplan
TEST?=$$(go list ./... |grep -v 'vendor')

all: plan

vet:
	go vet ./...

test: vet
	go test $(TEST)

build: test
	go install

init: build
	terraform init -plugin-dir $(GOPATH)bin

plan: init
	terraform plan -out ${TFPLAN}

acc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

apply:
	terraform apply ${TFPLAN}