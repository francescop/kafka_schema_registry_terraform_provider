all:
	go build -o terraform-provider-schemaregistry
	terraform init
	terraform apply -auto-approve
