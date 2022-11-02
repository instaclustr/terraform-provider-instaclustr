---
page_title: "instaclustr_iam_principal_arns_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_iam_principal_arns_v2_instance (Data Source)
Definition of an IAM Principal ARN being used for Kafka PrivateLink.
## Example Usage
```
data "instaclustr_iam_principal_arns_v2_instance" "example" { 
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
*___principal_arn___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^arn:aws:iam::[0-9]{12}:(root$|user\/[\w+=,.@-]+|role\/[\w+=,.@-]+)$`<br><br>The IAM Principal ARN.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the IAM Principal ARN.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the cluster data center.<br><br>
*___hidden_property_ignore___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
