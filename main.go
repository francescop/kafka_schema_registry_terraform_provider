package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
  "github.com/francescop/terraform_provider_kafka_schema_registry/restapi"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() terraform.ResourceProvider {
      return restapi.Provider()
    },
  })
}
