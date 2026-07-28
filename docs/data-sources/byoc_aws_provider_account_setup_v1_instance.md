---
page_title: "instaclustr_byoc_aws_provider_account_setup_v1_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_byoc_aws_provider_account_setup_v1_instance (Data Source)
A resource representing a BYOC provider account setup record.
## Example Usage
```
data "instaclustr_byoc_aws_provider_account_setup_v1_instance" "example" { 
  provider_account_id = "<provider_account_id>" // the value of the `provider_account_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
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
*___iam_role_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Custom IAM role name for BYOC setup. If not provided, defaults to 'instaclustr-<account-id>'. Must be 1-64 characters, containing only alphanumeric characters and +=,.@-_ symbols.<br><br>
*___cloud_provider_account_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Cloud provider account ID for BYOC setup.<br><br>
*___provider_account_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the provider account.<br><br>
*___data_centre_backup_buckets___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [data_centre_backup_buckets](#nested--data_centre_backup_buckets) for nested schema<br>
<br>Ordered list of data centre backup bucket mappings. The first entry is the primary region. Exactly one entry is required on create.<br><br>
*___skip_key_pair_creation___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>When true, an EC2 key pair already exists for this Instaclustr account in the target AWS account and primary data centre region (created by another RUNNING BYOC setup, or by this setup if already validated). Terraform users should pass create_key_pair = false when applying the template. CloudFormation downloads set CreateKeyPair automatically.<br><br>
*___provider_account_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the provider account.<br><br>
<a id="nested--data_centre_backup_buckets"></a>
## Nested schema for `data_centre_backup_buckets`
Ordered list of data centre backup bucket mappings. The first entry is the primary region. Exactly one entry is required on create.<br>
### Read-only attributes
*___data_centre___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Data centre region name.<br><br>
*___backup_bucket_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>S3 backup bucket name for the data centre region.<br><br>
