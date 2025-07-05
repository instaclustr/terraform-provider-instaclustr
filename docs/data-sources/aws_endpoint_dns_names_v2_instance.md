---
page_title: "instaclustr_aws_endpoint_dns_names_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_endpoint_dns_names_v2_instance (Data Source)
List of AWS endpoint DNS names for a cluster data center
## Example Usage
```
data "instaclustr_aws_endpoint_dns_names_v2_instance" "example" { 
  cluster_data_center_id = "<cluster_data_center_id>" // the value of the `cluster_data_center_id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.
- **_required_** - These attributes must be provided for the data source's information to be queried.
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
## Root Level Schema
### Read-only attributes
*___aws_endpoint_dns_names___*<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>AWS Endpoint DNS Names.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the cluster data center<br><br>
