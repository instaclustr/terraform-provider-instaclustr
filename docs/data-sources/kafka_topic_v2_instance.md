---
page_title: "instaclustr_kafka_topic_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_topic_v2_instance (Data Source)
Definition of a Kafka Topic to be applied to a Kafka cluster.
## Example Usage
```
data "instaclustr_kafka_topic_v2_instance" "example" { 
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
*___replication_factor___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Replication factor for Topic<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Kafka topic. The value of this property has the form: [cluster-id]_[kafka-topic]<br><br>
*___partitions___*<br>
<ins>Type</ins>: integer, read-only<br>
<br>Topic partition count<br><br>
*___topic___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Kafka Topic name<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Kafka cluster<br><br>
*___configs___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [configs](#nested--configs) for nested schema<br>
<br>List of Kafka topic configs which have non-default values. These could be set by terraform or other methods like kafka cli etc.<br><br>
<a id="nested--configs"></a>
## Nested schema for `configs`
List of Kafka topic configs which have non-default values. These could be set by terraform or other methods like kafka cli etc.<br>
### Read-only attributes
*___key___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Kafka Topic config key<br><br>
*___value___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Kafka Topic config value<br><br>
