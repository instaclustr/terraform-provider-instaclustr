---
page_title: "instaclustr_kafka_user_v5_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_user_v5_instance (Data Source)
Definition of a Kafka User to be applied to a Kafka cluster.
## Example Usage
```
data "instaclustr_kafka_user_v5_instance" "example" { 
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
*___auth_mechanism___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `MTLS`, `SASL` ]<br><br>Authentication mechanisms supported for KafkaClusters.<br><br>
*___username___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^(?![zZ][oO][oO][kK][eE][eE][pP][eE][rR]$)[a-zA-Z0-9][a-zA-Z0-9_-]*$`<br><br>Username of the Kafka user.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Kafka user. The value of this property has the form: [cluster-id]_[kafka-username]<br><br>
*___current_operation_status___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: allowed values: [ `NO_OPERATION`, `OPERATION_IN_PROGRESS`, `OPERATION_FAILED` ]<br><br>Indicates if the cluster is currently performing any operation such as being created, updated, or deleted<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Kafka cluster.<br><br>
*___sasl_scram_mechanism___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Scram Mechanism for SASL authentication. Valid values: SCRAM-SHA-256, SCRAM-SHA-512<br><br>
*___password___*<br>
<ins>Type</ins>: string (password), read-only<br>
<br>Password for the Kafka user.<br><br>
