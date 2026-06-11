---
page_title: "instaclustr_opensearch_external_s3_bucket_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_opensearch_external_s3_bucket_v2 (Resource)
Defines an external S3 bucket attached to an OpenSearch cluster for read-only access
## Example Usage
```
resource "instaclustr_opensearch_external_s3_bucket_v2" "example" {
  s3_bucket_name = "my-snapshot-bucket"
  cluster_id = "71e4380e-32ac-4fa7-ab42-c165fe35aa55"
  s3_prefixes = [ "snapshots/my-cluster" ]
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
*___s3_bucket_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Name of the external S3 bucket<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the OpenSearch cluster to attach the bucket to<br><br>
### Input attributes - Optional
*___s3_prefixes___*<br>
<ins>Type</ins>: list of strings, optional, immutable<br>
<br>S3 path prefixes to restrict read-only access to within the bucket<br><br>
### Read-only attributes
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Instaclustr ID of the external S3 bucket configuration<br><br>
*___cluster_data_centre_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the AWS VPC cluster data centre the bucket permission is attached to<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_opensearch_external_s3_bucket_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
