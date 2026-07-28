---
page_title: "instaclustr_byoc_aws_provider_account_setup_v1 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_byoc_aws_provider_account_setup_v1 (Resource)
A resource representing a BYOC provider account setup record.
## Example Usage
```
resource "instaclustr_byoc_aws_provider_account_setup_v1" "example" {
  cloud_provider_account_id = "123456789012"
  data_centre_backup_buckets {
    backup_bucket_name = "my-byoc-backup-bucket-east"
    data_centre = "US_EAST_1"
  }

  provider_account_name = "my-byoc-provider-account"
  iam_role_name = "instaclustr-my-account-role"
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
*___cloud_provider_account_id___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Cloud provider account ID for BYOC setup.<br><br>
*___data_centre_backup_buckets___*<br>
<ins>Type</ins>: repeatable nested block, required, updatable, see [data_centre_backup_buckets](#nested--data_centre_backup_buckets) for nested schema<br>
<br>Ordered list of data centre backup bucket mappings. The first entry is the primary region. Exactly one entry is required on create.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Name of the provider account.<br><br>
### Input attributes - Optional
*___iam_role_name___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>Custom IAM role name for BYOC setup. If not provided, defaults to 'instaclustr-<account-id>'. Must be 1-64 characters, containing only alphanumeric characters and +=,.@-_ symbols.<br><br>
### Read-only attributes
*___cloud_provider___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Cloud provider for this BYOC setup.<br><br>
*___self_service_byoc___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Whether this provider account was created via self-service BYOC setup.<br><br>
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the BYOC provider account setup.<br><br>
*___existing_key_pair_provider_account_names___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>Names of other RUNNING BYOC provider accounts under this Instaclustr account that already share the same AWS account ID and primary data centre region. Empty when skipKeyPairCreation is false.<br><br>
*___account_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>UUID of the Instaclustr Account.<br><br>
*___provider_account_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the provider account.<br><br>
*___skip_key_pair_creation___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>When true, an EC2 key pair already exists for this Instaclustr account in the target AWS account and primary data centre region (created by another RUNNING BYOC setup, or by this setup if already validated). Terraform users should pass create_key_pair = false when applying the template. CloudFormation downloads set CreateKeyPair automatically.<br><br>
<a id="nested--data_centre_backup_buckets"></a>
## Nested schema for `data_centre_backup_buckets`
Ordered list of data centre backup bucket mappings. The first entry is the primary region. Exactly one entry is required on create.<br>
### Input attributes - Required
*___data_centre___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>Data centre region name.<br><br>
*___backup_bucket_name___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>S3 backup bucket name for the data centre region.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_byoc_aws_provider_account_setup_v1.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `provider_account_id` attribute defined in the root schema above.
