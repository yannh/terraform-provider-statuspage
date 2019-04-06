#!/usr/bin/make

TFPLAN ?= plan.tfplan
TEST?=$$(go list ./... |grep -v 'vendor')
export CGO_ENABLED = 0

all: plan

vet:
	go vet ./...

test: vet
	go test $(TEST)

build:
	go install

init: test build
	terraform init -plugin-dir $(GOPATH)bin

plan: init
	terraform plan -out ${TFPLAN}

acc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

apply:
	terraform apply ${TFPLAN}
