plan:
	terraform -chdir=terraform plan -var-file="local.tfvars" 

apply:
	terraform -chdir=terraform apply -var-file="local.tfvars"

destroy:
	terraform -chdir=terraform destroy -var-file="local.tfvars"