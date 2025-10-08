---
page_title: "instaclustr_aws_fsxn_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_fsxn_v2_instance (Data Source)
AWS FSx ONTAP file system.
## Example Usage
```
data "instaclustr_aws_fsxn_v2_instance" "example" { 
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
<br>Status of the file system<br><br>
*___vpc_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The ID of the VPC where this file system should be provisioned.<br><br>
*___fsxn_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>AWS ID of the file system.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Instaclustr ID representing the file system.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The ID of the cluster who's VPC this file system should share.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the provider account.<br><br>
