develop:
	mkdir -p ~/.terraform.d/plugins/theostanton/providers/notion/0.0.1/darwin_amd64
	rm -rf usage/.terraform
	rm -f usage/.terraform.lock.hcl
	go build -o terraform-provider-notion
	mv terraform-provider-notion ~/.terraform.d/plugins/theostanton/providers/notion/0.0.1/darwin_amd64
	cd usage && terraform init
	cd usage && TF_LOG=INFO terraform apply

destroy:
	cd usage && terraform destroy -auto-approve
	rm -f usage/terraform.tfstate
	rm -f usage/terraform.tfstate.backup