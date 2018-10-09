package restapi

import (
  "github.com/hashicorp/terraform/helper/schema"
  "log"
)

func resourceSubject() *schema.Resource {
  return &schema.Resource{
    Create: resourceSubjectCreate,
    Read:   resourceSubjectRead,
    Update: resourceSubjectUpdate,
    Delete: resourceSubjectDelete,

    Schema: map[string]*schema.Schema{
      "subject": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "schema": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
  }
}

func resourceSubjectCreate(d *schema.ResourceData, m interface{}) error {
  client, err := make_client(d, m)

  if err != nil { return err }

  log.Printf("Create subject '%s'.", client)

  err = client.create_subject()

  if err != nil {
    return err
  }

  d.SetId(client.subject)
  return nil
}

func resourceSubjectRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceSubjectUpdate(d *schema.ResourceData, m interface{}) error {
  client, err := make_client(d, m)

  if err != nil { return err }

  log.Printf("Update subject '%s'.", client)

  err = client.create_subject()

  if err != nil {
    return err
  }

  return nil
}

func resourceSubjectDelete(d *schema.ResourceData, m interface{}) error {
  client, err := make_client(d, m)

  if err != nil { return err }

  log.Printf("Delete subject '%s'.", client)

  err = client.delete_subject()

  if err != nil {
    return err
  }

  d.SetId("")

  return nil
}
