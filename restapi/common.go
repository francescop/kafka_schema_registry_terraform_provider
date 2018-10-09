package restapi

import (
  "github.com/hashicorp/terraform/helper/schema"
)

func make_client(d *schema.ResourceData, m interface{}) (*schema_registry_client, error) {
  uri     := m.(string)
  subject := d.Get("subject").(string)
  schema  := d.Get("schema").(string)

  client, err := NewSchemaRegistryClient(uri, subject, schema)

  return client, err
}
