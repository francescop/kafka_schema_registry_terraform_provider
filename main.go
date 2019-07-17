package main

import (
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
  "github.com/francescop/kafka_schema_registry_terraform_provider/restapi"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() terraform.ResourceProvider {
      return restapi.Provider()
    },
  })
}
