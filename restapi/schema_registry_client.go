package restapi

import (
  "encoding/json"
  "bytes"
  "io/ioutil"
  "net/http"
  "fmt"
)

type schema_registry_client struct {
  http_client           *http.Client
  create_uri            string
  update_uri            string
  read_uri              string
  delete_uri            string
  subject               string
  schema                string
}

func NewSchemaRegistryClient(uri string, subject string, schema string) (*schema_registry_client, error) {
  client := schema_registry_client{
    create_uri: uri + "/subjects/" + subject + "/versions",
    update_uri: uri + "/subjects/" + subject + "/versions",
    read_uri: uri + "/subjects/" + subject + "/versions",
    delete_uri: uri + "/subjects/" + subject,
    subject: subject,
    schema: schema,
    http_client: &http.Client{},
  }

  return &client, nil
}

func (client schema_registry_client) create_subject() error {
  jsonData := map[string]string{"schema": client.schema}
  jsonValue, err := json.Marshal(jsonData)

  if err != nil {
    return err
  }

  response, err := http.Post(client.create_uri, "application/vnd.schemaregistry.v1+json", bytes.NewBuffer(jsonValue))

  if err != nil {
    return err
  }

  data, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return err
  }

  if response.StatusCode != http.StatusOK {
    err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
    return err
  }

  return nil
}

func (client schema_registry_client) delete_subject() error {
  request, err := http.NewRequest("DELETE", client.delete_uri, nil)

  if err != nil {
    return err
  }

  response, err := client.http_client.Do(request)

  if err != nil {
    return err
  }

  data, err := ioutil.ReadAll(response.Body)

  if err != nil {
    return err
  }

  if response.StatusCode != http.StatusOK {
    err = fmt.Errorf("response code is %d: %s", response.StatusCode, data)
    return err
  }

  defer response.Body.Close()

  return err
}
