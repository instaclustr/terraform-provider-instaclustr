---
page_title: "instaclustr_kafka_cluster_kafka_topics_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_cluster_kafka_topics_v2 (Data Source)
Definition of Kafka Topic names in the Kafka cluster
## Example Usage
```
data "instaclustr_kafka_cluster_kafka_topics_v2" "example" { 
  kafka_cluster_id = "<kafka_cluster_id>" // the value of the `kafka_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___kafka_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the Kafka cluster<br><br>
### Read-only attributes
*___topic_names___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>List of Kafka topic names<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Kafka cluster<br><br>
