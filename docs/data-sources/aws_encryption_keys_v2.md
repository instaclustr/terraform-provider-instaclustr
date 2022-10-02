---
page_title: "instaclustr_aws_encryption_keys_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_encryption_keys_v2 (Data Source)
A listable data source of all AWS Encryption Keys in an Instaclustr Account.
## Example Usage
```
data "instaclustr_aws_encryption_keys_v2" "example" { }
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.<br><br>
- **_required_** - These attributes must be provided for the data source's information to be queried.<br><br>
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
# Schema
## Read-only attributes
### account_id<br>
<ins>Type</ins>: string, read-only<br>
<br>UUID of the Instaclustr Account.
### encryption_keys<br>
<ins>Type</ins>: repeatable nested block, read-only, see [encryption_keys](#nested--encryption_keys) for nested schema<br>

<a id="nested--encryption_keys"></a>
# Nested schema for `encryption_keys`<br>

## Read-only attributes
### in_use<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether the encryption key is used by a cluster.
### id<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the encryption key.
### arn<br>
<ins>Type</ins>: string, read-only<br>
<br>AWS ARN for the encryption key.
### alias<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]{1}[a-zA-Z0-9 _-]*$`<br><br>Encryption key alias for display purposes.
### provider_account_name<br>
<ins>Type</ins>: string, read-only<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.
