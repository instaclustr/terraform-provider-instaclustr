---
page_title: "instaclustr_aws_endpoint_dns_names_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_endpoint_dns_names_v2 (Resource)
List of AWS endpoint DNS names for a cluster data center
## Example Usage
```
resource "instaclustr_aws_endpoint_dns_names_v2" "example" {
  cluster_data_center_id = "1e4bd709-c4ed-43f8-b5fa-5964d3ef3c1e"
  aws_endpoint_dns_names = [ "ip-10-0-169-78.ec2.internal", "ip-10-0-107-16.ec2.internal", "ip-10-0-52-211.ec2.internal" ]
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
*___aws_endpoint_dns_names___*<br>
<ins>Type</ins>: list of strings, required, updatable<br>
<br>AWS Endpoint DNS Names.<br><br>
*___cluster_data_center_id___*<br>
<ins>Type</ins>: string, required, updatable<br>
<br>ID of the cluster data center<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_endpoint_dns_names_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `cluster_data_center_id` attribute defined in the root schema above.
