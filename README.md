# Terraform provider for Confluent Kafka Schema Registry

This terraform provider allows you to interact with Confluent Kafka Schema Registry.

There are a few requirements about how the API must work for this provider to be able to do its thing:
* The API is expected to support the following HTTP methods:
    * POST: create or update a schema
    * DELETE: remove a schema

Have a look at the [examples directory](examples) for some use cases

&nbsp;

## Provider configuration
- `uri` (string, required): URI of the Schema Registry REST API endpoint. This serves as the base of all requests. Example: `http://localhost:8001`.

&nbsp;

## `schemaregistry_subject` resource configuration
- `subject` (string, required): The name of the subject to be created.
- `schema` (string, required): The schema of the subject that has to be created.

&nbsp;

&nbsp;

## Installation
There are two standard methods of installing this provider detailed [in Terraform's documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins). You can place the file in the directory of your .tf file in `terraform.d/plugins/{OS}_{ARCH}/` or place it in your home directory at `~/.terraform.d/plugins/{OS}_{ARCH}/`

Once downloaded, be sure to make the plugin executable by running `chmod +x terraform-provider-schemaregistry`.

## Extras

In the [examples directory](examples) there is a `docker-compose.yml` file.
This is useful for development or for testing purposes.

## Todo

[x] write docker compose file
[x] write samples
[] write tests
