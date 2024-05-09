---
page_title: "instaclustr_kafka_user_certificate_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_kafka_user_certificate_v2_instance (Data Source)
Certificate signing request.
## Example Usage
```
data "instaclustr_kafka_user_certificate_v2_instance" "example" { 
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
*___kafka_username___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The Kafka username<br><br>
*___valid_period___*<br>
<ins>Type</ins>: integer, read-only<br>
<ins>Constraints</ins>: minimum: 3, maximum: 1.2E+2<br><br>Number of months for which the certificate will be valid.<br><br>
*___id___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>ID of the certificate.<br><br>
*___csr___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>Certificate signing request.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the kafka cluster<br><br>
*___expiry_date___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>Date certificate expires.<br><br>
*___signed_certificate___*<br>
<ins>Type</ins>: string (string), read-only<br>
<br>Generated client signed certificate.<br><br>
