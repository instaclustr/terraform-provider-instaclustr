---
page_title: "instaclustr_clickhouse_integration_private_s3_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_integration_private_s3_v2 (Resource)
ClickHouse integration with Private S3 buckets - enables access.
## Example Usage
```
resource "instaclustr_clickhouse_integration_private_s3_v2" "example" {
  kms_key_arns = [ "arn:aws:kms:us-west-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab" ]
  buckets {
    arn = "arn:aws:s3:::example-s3-bucket"
    read_only = false
  }

  buckets {
    arn = "arn:aws:s3:::example-s3-bucket2"
    read_only = true
  }

  cluster_id = "71e4380e-32ac-4fa7-ab42-c165fe35aa55"
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
*___buckets___*<br>
<ins>Type</ins>: list of objects, required, updatable<br>
<ins>Constraints</ins>: minimum items: 1<br><br>S3 buckets for the Private S3 Integration.<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>ID of the ClickHouse cluster<br><br>
### Input attributes - Optional
*___kms_key_arns___*<br>
<ins>Type</ins>: list of strings, optional, updatable<br>
<ins>Constraints</ins>: minimum items: 0<br><br>List of KMS key ARNs.<br><br>
### Read-only attributes
*___status___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the Private S3 integration.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Private S3 Integration.<br><br>
<a id="nested--buckets"></a>
## Nested schema for `buckets`
S3 buckets for the Private S3 Integration.<br>
### Input attributes - Optional
*___read_only___*<br>
<ins>Type</ins>: boolean, optional, updatable<br>
<br>Boolean recording whether the S3 bucket is read-only.<br><br>
*___arn___*<br>
<ins>Type</ins>: string, optional, updatable<br>
<br>ARN for the S3 bucket<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_clickhouse_integration_private_s3_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
