---
page_title: "instaclustr_clickhouse_integration_private_s3_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_clickhouse_integration_private_s3_v2_instance (Data Source)
ClickHouse Integration With Private S3 Buckets - Enables Access
## Example Usage
```
data "instaclustr_clickhouse_integration_private_s3_v2_instance" "example" { 
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
<br>Status of the private S3 Integration<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the Private S3 Integration<br><br>
*___buckets___*<br>
<ins>Type</ins>: list of objects, read-only<br>
<ins>Constraints</ins>: minimum items: 1<br><br>S3 buckets for the Private S3 Integration<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the ClickHouse cluster<br><br>
*___kms_key_arns___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<ins>Constraints</ins>: minimum items: 0<br><br>List of KMS key ARNs<br><br>
<a id="nested--buckets"></a>
## Nested schema for `buckets`
S3 buckets for the Private S3 Integration<br>
### Read-only attributes
*___read_only___*<br>
<ins>Type</ins>: boolean, read-only<br>
<br>Boolean recording whether the S3 bucket is read-only<br><br>
*___arn___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ARN for the S3 bucket<br><br>
