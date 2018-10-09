package restapi

import (
  "github.com/hashicorp/terraform/helper/schema"
  "net/url"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "uri": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        Description: "Kafka schema registry endpoint. Example: http://localhost:8000",
        },
      },
    ResourcesMap: map[string]*schema.Resource{
      "schemaregistry_subject": resourceSubject(),
    },
    ConfigureFunc: configureProvider,
  }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
  endpoint := d.Get("uri").(string)
  _, err := url.ParseRequestURI(endpoint)

  return endpoint, err
}
