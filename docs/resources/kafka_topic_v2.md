---
page_title: "instaclustr_kafka_topic_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_topic_v2 (Resource)
Definition of a Kafka Topic to be applied to a Kafka cluster.
## Example Usage
```
resource "instaclustr_kafka_topic_v2" "example" {
  partitions = 3
  replication_factor = 3
  topic = "topic-test"
  cluster_id = "c1af59c6-ba0e-4cc2-a0f3-65cee17a5f37"
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
*___replication_factor___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<br>Replication factor for Topic<br><br>
*___partitions___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<br>Topic partition count<br><br>
*___topic___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Kafka Topic name<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Kafka cluster<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Kafka topic. The value of this property has the form: [cluster-id]_[kafka-topic]<br><br>
*___configs___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [configs](#nested--configs) for nested schema<br>
<br>List of the the Kafka cluster configs<br><br>
<a id="nested--configs"></a>
## Nested schema for `configs`
List of the the Kafka cluster configs<br>
### Input attributes - Required
*___key___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Kafka Topic config key<br><br>
*___value___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Kafka Topic config value<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_topic_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
