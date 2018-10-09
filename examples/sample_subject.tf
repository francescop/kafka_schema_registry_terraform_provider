provider "schemaregistry" {
  uri = "http://localhost:8081"
}

resource "schemaregistry_subject" "schema_sample_from_string" {
  subject = "com.test.myapp.test-from-string"
  schema  = "{\"type\": \"string\"}"
}

resource "schemaregistry_subject" "schema_sample_from_file" {
  subject = "com.test.myapp.test-from-file"
  schema  = "${file("schema_sample.avro.json")}"
}
