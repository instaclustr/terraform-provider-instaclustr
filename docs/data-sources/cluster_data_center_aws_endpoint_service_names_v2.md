---
page_title: "instaclustr_cluster_data_center_aws_endpoint_service_names_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_cluster_data_center_aws_endpoint_service_names_v2 (Data Source)
List of AWS endpoint service names for a cluster data center
## Example Usage
```
data "instaclustr_cluster_data_center_aws_endpoint_service_names_v2" "example" { 
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
*___aws_endpoint_service_names___*<br>
<ins>Type</ins>: repeatable nested block, read-only, see [aws_endpoint_service_names](#nested--aws_endpoint_service_names) for nested schema<br>
<br>AWS Endpoint Service Names.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster data center<br><br>
<a id="nested--aws_endpoint_service_names"></a>
## Nested schema for `aws_endpoint_service_names`
AWS Endpoint Service Names.<br>
### Read-only attributes
*___end_point_service_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The Instaclustr ID of the AWS endpoint service<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>The ID of the cluster data center.<br><br>
*___end_point_service_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The Endpoint Service Name.<br><br>
