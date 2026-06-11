---
page_title: "instaclustr_opensearch_cluster_opensearch_external_s3_bucket_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_opensearch_cluster_opensearch_external_s3_bucket_v2 (Data Source)
List of external S3 bucket configurations with read-only access for an OpenSearch cluster
## Example Usage
```
data "instaclustr_opensearch_cluster_opensearch_external_s3_bucket_v2" "example" { 
  opensearch_cluster_id = "<opensearch_cluster_id>" // the value of the `opensearch_cluster_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___opensearch_cluster_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the OpenSearch cluster<br><br>
### Read-only attributes
*___external_s3_buckets___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [external_s3_buckets](#nested--external_s3_buckets) for nested schema<br>
<br>List of external S3 bucket configurations<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the OpenSearch cluster<br><br>
<a id="nested--external_s3_buckets"></a>
## Nested schema for `external_s3_buckets`
List of external S3 bucket configurations<br>
### Read-only attributes
*___s3_bucket_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the external S3 bucket<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>Instaclustr ID of the external S3 bucket configuration<br><br>
*___cluster_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the OpenSearch cluster to attach the bucket to<br><br>
*___s3_prefixes___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>S3 path prefixes to restrict read-only access to within the bucket<br><br>
*___cluster_data_centre_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the AWS VPC cluster data centre the bucket permission is attached to<br><br>
