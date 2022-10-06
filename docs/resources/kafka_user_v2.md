---
page_title: "instaclustr_kafka_user_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_user_v2 (Resource)
Definition of a Kafka User to be applied to a Kafka cluster.
## Example Usage
```
resource "instaclustr_kafka_user_v2" "example" {
  password = "myPassword1."
  options = {
    override_existing_user = false
    sasl_scram_mechanism = "SCRAM-SHA-256"
  }

  cluster_id = "c1af59c6-ba0e-4cc2-a0f3-65cee17a5f37"
  initial_permissions = "standard"
  username = "myKafkaUser"
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
*___options___*<br>
<ins>Type</ins>: nested object, required, updatable, see [options](#nested--options) for nested schema<br>
<br>Initial options used when creating Kafka user<br><br>
*___username___*<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `^(?![zZ][oO][oO][kK][eE][eE][pP][eE][rR]$)[a-zA-Z0-9][a-zA-Z0-9_-]*$`<br><br>Username of the Kafka user.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Kafka cluster.<br><br>
*___password___*<br>
<ins>Type</ins>: string (password), required, updatable<br>
<br>Password for the Kafka user.<br><br>
### Input attributes - Optional
*___initial_permissions___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<ins>Constraints</ins>: allowed values: [ `standard`, `read-only`, `none` ]<br><br>Permissions initially granted to Kafka user upon creation.<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr identifier for the Kafka user. The value of this property has the form: [cluster-id]_[kafka-username]<br><br>
<a id="nested--options"></a>
## Nested schema for `options`
Initial options used when creating Kafka user<br>
### Input attributes - Required
*___sasl_scram_mechanism___*<br>
<ins>Type</ins>: string, required, updatable<br>
<ins>Constraints</ins>: allowed values: [ `SCRAM-SHA-256`, `SCRAM-SHA-512` ]<br><br>SASL/SCRAM mechanism for user<br><br>
### Input attributes - Optional
*___override_existing_user___*<br>
<ins>Type</ins>: boolean, optional, immutable<br>
<br>Overwrite user if already exists.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_user_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
