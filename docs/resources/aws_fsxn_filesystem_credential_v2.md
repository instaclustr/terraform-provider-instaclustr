---
page_title: "instaclustr_aws_fsxn_filesystem_credential_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_fsxn_filesystem_credential_v2 (Resource)
Response body for FSxN admin credential submission.
## Example Usage
```
resource "instaclustr_aws_fsxn_filesystem_credential_v2" "example" {
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
### Input attributes - Optional
*___username___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Username for the filesystem-level administrative user (fsxadmin role).<br><br>
*___password___*<br>
<ins>Type</ins>: string (password), optional, immutable<br>
<br>Password for the filesystem-level administrative user.<br><br>
### Read-only attributes
*___filesystem_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The AWS FSxN filesystem ID (e.g. fs-0c315066d18074b67).<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_fsxn_filesystem_credential_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `filesystem_id` attribute defined in the root schema above.
