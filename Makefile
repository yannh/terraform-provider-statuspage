#!/usr/bin/make

TFPLAN ?= plan.tfplan
GOIMAGE ?= golang:1.18.1

export CGO_ENABLED = 0
export GOFLAGS = -mod=vendor
export GOOS = linux
export GOARCH = amd64

all: plan

vet:
	go vet $(TEST)

test: vet
	go test -v $(TEST)

docker-test:
	docker run -t -v $$PWD:/go/src/github.com/yannh/terraform-provider-statuspage -w /go/src/github.com/yannh/terraform-provider-statuspage $(GOIMAGE) make test

build build-static:
	go build -trimpath -tags netgo -ldflags '-extldflags "-static"' -a -o bin/terraform-provider-statuspage

docker-build-static:
	docker run -t -v $$PWD:/go/src/github.com/yannh/terraform-provider-statuspage -w /go/src/github.com/yannh/terraform-provider-statuspage $(GOIMAGE) make build-static

goreleaser-build-static:
	docker run -e GOCACHE=/tmp -v $$PWD/.gitconfig:/root/.gitconfig -t -v $$PWD:/go/src/github.com/yannh/terraform-provider-statuspage -w /go/src/github.com/yannh/terraform-provider-statuspage goreleaser/goreleaser:v1.8.3 build --single-target --skip-post-hooks --rm-dist --snapshot

init: test build-static
	terraform init -plugin-dir ./bin

plan: init
	terraform plan -out ${TFPLAN}

acc:
	TF_ACC=1 go test ./... -v -timeout 120m -count=1

docker-acc:
	docker run -e DATADOG_API_KEY -e DATADOG_APPLICATION_KEY -e STATUSPAGE_PAGE -e STATUSPAGE_PAGE_2 -e STATUSPAGE_TOKEN -t -v $$PWD:/go/src/github.com/yannh/terraform-provider-statuspage -w /go/src/github.com/yannh/terraform-provider-statuspage $(GOIMAGE) make acc

apply:
	terraform apply ${TFPLAN}

update-sdk:
	GOFLAGS=  go get -u github.com/yannh/statuspage-go-sdk
	go mod vendor

release:
	docker run -e GITHUB_TOKEN -t -v $$PWD/.gitconfig:/root/.gitconfig -v /var/run/docker.sock:/var/run/docker.sock -v $$PWD:/go/src/github.com/yannh/terraform-provider-statuspage -w /go/src/github.com/yannh/terraform-provider-statuspage goreleaser/goreleaser:v1.8.3 release --rm-dist

