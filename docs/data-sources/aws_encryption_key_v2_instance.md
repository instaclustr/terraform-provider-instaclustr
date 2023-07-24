---
page_title: "instaclustr_aws_encryption_key_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_encryption_key_v2_instance (Data Source)
Definition of a customer managed AWS KMS Key for use with at rest EBS encryption in Instaclustr managed clusters.
## Example Usage
```
data "instaclustr_aws_encryption_key_v2_instance" "example" { 
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
*___in_use___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether the encryption key is used by a cluster.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the encryption key.<br><br>
*___arn___*<br>
<ins>Type</ins>: string, read-only<br>
<br>AWS ARN for the encryption key.<br><br>
*___alias___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]{1}[a-zA-Z0-9 _-]*$`<br><br>Encryption key alias for display purposes.<br><br>
*___regions___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [regions](#nested--regions) for nested schema<br>
<br>Regions of the key<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.<br><br>
<a id="nested--regions"></a>
## Nested schema for `regions`
Regions of the key<br>
### Read-only attributes
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the AWS Region<br><br>
