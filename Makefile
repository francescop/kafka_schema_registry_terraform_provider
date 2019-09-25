openbsd_amd64:
	GOOS=openbsd GOARCH=amd64 go build -o terraform-provider-schemaregistry_openbsd_amd64

freebsd_amd64:
	GOOS=freebsd GOARCH=amd64 go build -o terraform-provider-schemaregistry_freebsd_amd64

linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-schemaregistry_linux_amd64

darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-schemaregistry-darwin_amd64

test:
	go build -o terraform-provider-schemaregistry
	cp terraform-provider-schemaregistry ~/.terraform.d/plugins/
	cd examples; terraform init; terraform apply -auto-approve

clean:
	rm -f terraform-provider-schemaregistry_*
	cd examples; rm -rf .terraform; rm -f *.tfstate*
