#!/usr/bin/make

build:
	go build -o terraform-provider-statuspage ./*.go

init: build
	terraform init

plan: init
	terraform plan -out plan.tfplan