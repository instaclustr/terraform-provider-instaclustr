---
page_title: "instaclustr_cluster_data_center_iam_principal_arns_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_data_center_iam_principal_arns_v2 (Data Source)
List of IAM Principal ARNs for a cluster data center
## Example Usage
```
data "instaclustr_cluster_data_center_iam_principal_arns_v2" "example" { 
  cluster_data_center_id = "<cluster_data_center_id>" // the value of the `cluster_data_center_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Input attributes - Required
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string, required<br>
<br>ID of the Cluster Data Center.<br><br>
### Read-only attributes
*___iam_principal_arns___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [iam_principal_arns](#nested--iam_principal_arns) for nested schema<br>
<br>IAM Principal ARNs to allow connection to the AWS Endpoint Service.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster data center<br><br>
<a id="nested--iam_principal_arns"></a>
## Nested schema for `iam_principal_arns`
IAM Principal ARNs to allow connection to the AWS Endpoint Service.<br>
### Read-only attributes
*___principal_arn___*<br>
<ins>Type</ins>: string, read-only<br>
<ins>Constraints</ins>: pattern: `^arn:aws:iam::[0-9]{12}:(root$|user\/[\w+=,.@-]+|role\/[\w+=,.@-]+)$`<br><br>The IAM Principal ARN.<br><br>
*___id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the IAM Principal ARN.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The Instaclustr ID of the cluster data center.<br><br>
*___hidden_property_ignore___*<br>
<ins>Type</ins>: string, read-only<br>
<br>
