---
page_title: "instaclustr_aws_encryption_key_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_encryption_key_v2 (Resource)
Definition of a customer managed AWS KMS Key for use with at rest EBS encryption in Instaclustr managed clusters.
## Example Usage
```
resource "instaclustr_aws_encryption_key_v2" "example" {
  alias = "encryption_key"
  arn = "arn:aws:kms:us-east-1:123456789123:key/123abcde-4567-8910-abcd-123456789abc"
}
```
## Glossary
The following terms are used to describe attributes in the schema of this resource:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the resource.<br><br>
- **_required_** - These attributes must be provided for the resource to be created.<br><br>
- **_optional_** - These input attributes can be omitted, and doing so may result in a default value being used.<br><br>
- **_immutable_** - These are input attributes that cannot be changed after the resource is created.<br><br>
- **_updatable_** - These input attributes can be updated to a different value if needed, and doing so will trigger an update operation.<br><br>
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
# Schema
## Input attributes - Required
### arn<br>
<ins>Type</ins>: string, required, immutable<br>
<br>AWS ARN for the encryption key.
### alias<br>
<ins>Type</ins>: string, required, immutable<br>
<ins>Constraints</ins>: pattern: `^[a-zA-Z0-9_-]{1}[a-zA-Z0-9 _-]*$`<br><br>Encryption key alias for display purposes.
## Input attributes - Optional
### provider_account_name<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>For customers running in their own account. Your provider account can be found on the Create Cluster page on the Instaclustr Console, or the "Provider Account" property on any existing cluster. For customers provisioning on Instaclustr's cloud provider accounts, this property may be omitted.
## Read-only attributes
### in_use<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether the encryption key is used by a cluster.
### id<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the encryption key.
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_encryption_key_v2.<resource-name> "<resource-id>"
```
`<resource-id>` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
