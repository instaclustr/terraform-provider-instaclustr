---
page_title: "instaclustr_aws_vpc_peer_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_vpc_peer_v2 (Resource)
Definition of an AWS VPC Peering request to allow privately routed connections to a target data centre.
## Example Usage
```
resource "instaclustr_aws_vpc_peer_v2" "example" {
  peer_aws_account_id = "123456789123"
  peer_subnets = [ "10.129.0.0/16" ]
  peer_vpc_id = "vpc-123abc456"
  peer_region = "US_EAST_1"
  cdc_id = "f3eab841-6952-430d-ba90-1bfc3f15da10"
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
*___peer_subnets___*<br>
<ins>Type</ins>: list of strings, required, updatable<br>
<br>The subnets for the peering VPC.<br><br>
*___cdc_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cluster Data Centre<br><br>
*___peer_region___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region code for the accepter VPC.<br><br>
*___peer_vpc_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>ID of the VPC with which the peering connection is created.<br><br>
*___peer_aws_account_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The AWS account ID of the owner of the accepter VPC.<br><br>
### Read-only attributes
*___status_code___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `pending-acceptance`, `failed`, `expired`, `provisioning`, `active`, `deleting`, `deleted` or `rejected`.<br><br>
*___data_centre_vpc_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the current data centre VPC.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_vpc_peer_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
