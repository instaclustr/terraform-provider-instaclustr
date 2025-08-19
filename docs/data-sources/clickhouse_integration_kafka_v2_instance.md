---
page_title: "instaclustr_clickhouse_integration_kafka_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_integration_kafka_v2_instance (Data Source)
ClickHouse Integration with Kafka Clusters - Enables Access
## Example Usage
```
data "instaclustr_clickhouse_integration_kafka_v2_instance" "example" { 
  id = "<id>" // the value of the `id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Kafka Integration<br><br>
*___kafka_cluster___*<br>
<ins>Type</ins>: object, read-only<br>
<br>Kafka cluster for the Kafka Integration<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Kafka Integration.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the ClickHouse cluster<br><br>
*___named_collection___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of of named collection used for the Kafka table integration config. Format: kafka-cluster-<integration-id>_topic_<topic-name><br><br>
