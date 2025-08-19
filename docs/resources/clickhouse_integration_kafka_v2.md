---
page_title: "instaclustr_clickhouse_integration_kafka_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_integration_kafka_v2 (Resource)
ClickHouse Integration with Kafka Clusters - Enables Access
## Example Usage
```
resource "instaclustr_clickhouse_integration_kafka_v2" "example" {
  cluster_id = "12aq0751-a169-8732-287a-21aedc6d23fd"
  kafka_cluster {
    format = "Avro"
    topic_name = "kafka-topic"
    id = "12aq0751-a169-8732-287a-21aedc6d23fc"
  }

}
```
## Glossary
The following terms are used to describe attributes in the schema of this resource:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the resource.
- **_required_** - These attributes must be provided for the resource to be created.
- **_optional_** - These input attributes can be omitted, and doing so may result in a default value being used.
- **_immutable_** - These are input attributes that cannot be changed after the resource is created.
- **_updatable_** - These input attributes can be updated to a different value if needed, and doing so will trigger an update operation.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___kafka_cluster___*<br>
<ins>Type</ins>: object, required, updatable<br>
<br>Kafka cluster for the Kafka Integration<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>ID of the ClickHouse cluster<br><br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Kafka Integration<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Kafka Integration.<br><br>
*___named_collection___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of of named collection used for the Kafka table integration config. Format: kafka-cluster-<integration-id>_topic_<topic-name><br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_clickhouse_integration_kafka_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
