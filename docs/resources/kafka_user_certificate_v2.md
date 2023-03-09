---
page_title: "instaclustr_kafka_user_certificate_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_user_certificate_v2 (Resource)
Certificate signing request.
## Example Usage
```
resource "instaclustr_kafka_user_certificate_v2" "example" {
  csr = "generated csr string"
  kafka_username = "kafka username"
  cluster_id = "Id of the kafka cluster"
  valid_period = 120
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
*___kafka_username___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The Kafka username<br><br>
*___valid_period___*<br>
<ins>Type</ins>: integer, required, immutable<br>
<ins>Constraints</ins>: minimum: 3, maximum: 1.2E+2<br><br>Number of months for which the certificate will be valid.<br><br>
*___csr___*<br>
<ins>Type</ins>: string (string), required, immutable<br>
<br>Certificate signing request.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>ID of the kafka cluster<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>ID of the certificate.<br><br>
*___expiry_date___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>Date certificate expires.<br><br>
*___signed_certificate___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>Generated client signed certificate.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_kafka_user_certificate_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
