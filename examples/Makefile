#!/usr/bin/make -f

plan: setup
	terraform plan

setup: clean
	mkdir -p terraform.d/plugins/registry.terraform.io/yannh/statuspage/0.1.10/linux_amd64/
	cp ../bin/terraform-provider-statuspage  terraform.d/plugins/registry.terraform.io/yannh/statuspage/0.1.10/linux_amd64/terraform-provider-statuspage
	terraform init -upgrade -plugin-dir terraform.d/plugins

clean:
	rm -rf .terraform .terraform.lock.hcl terraform.tfstate*

init:
	terraform init -plugin-dir terraform.d/plugins
