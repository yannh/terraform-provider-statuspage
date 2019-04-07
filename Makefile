#!/usr/bin/make

TFPLAN ?= plan.tfplan
TEST?=$$(go list ./... |grep -v 'vendor')
export CGO_ENABLED = 0

all: plan

vet:
	go vet $(TEST)

test: vet
	go test -mod=vendor $(TEST)

build:
	go install -mod=vendor -tags netgo -ldflags '-extldflags "-static"'

init: test build
	terraform init -plugin-dir $(GOPATH)bin

plan: init
	terraform plan -out ${TFPLAN}

acc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

apply:
	terraform apply ${TFPLAN}
