openbsd_amd64:
	GOOS=openbsd GOARCH=amd64 go build -o terraform-provider-schemaregistry
	tar cvzf terraform-provider-schemaregistry_openbsd_amd64.tar.gz terraform-provider-schemaregistry

freebsd_amd64:
	GOOS=freebsd GOARCH=amd64 go build -o terraform-provider-schemaregistry
	tar cvzf terraform-provider-schemaregistry_freebsd_amd64.tar.gz terraform-provider-schemaregistry

linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-schemaregistry
	tar cvzf terraform-provider-schemaregistry_linux_amd64.tar.gz terraform-provider-schemaregistry

darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-schemaregistry
	tar cvzf terraform-provider-schemaregistry_darwin_amd64.tar.gz terraform-provider-schemaregistry

test:
	go build -o terraform-provider-schemaregistry
	cp terraform-provider-schemaregistry ~/.terraform.d/plugins/
	cd examples; terraform init; terraform apply -auto-approve

clean:
	cd examples; rm -rf .terraform; rm -f *.tfstate*
