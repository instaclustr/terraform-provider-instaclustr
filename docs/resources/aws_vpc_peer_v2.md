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
The following terms are used to describe properties in the schema of this resource:
- **_read-only_** - These are properties that can only be read and not provided as an input to the resource.<br><br>
- **_required_** - These properties must be provided for the resource to be created.<br><br>
- **_optional_** - These input properties can be omitted, and doing so may result in a default value being used.<br><br>
- **_immutable_** - These are input properties that cannot be changed after the resource is created. The resource will be destroyed and re-created on `terraform apply` if Terraform detects a change in such properties.<br><br>
- **_updatable_** - These input properties can be updated to a different value if needed, and doing so will trigger an update operation.<br><br>
- **_nested block_** - These properties use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Properties with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested properties. When reading nested block properties, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_property[0].nested_property`.
## Schema
### peer_aws_account_id<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The AWS account ID of the owner of the accepter VPC.
### peer_subnets<br>
<ins>Type</ins>: list of arrays, required, updatable<br>
<br>The subnets for the peering VPC.
### peer_vpc_id<br>
<ins>Type</ins>: string, required, immutable<br>
<br>ID of the VPC with which the peering connection is created.
### id<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.
### peer_region<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Region code for the accepter VPC.
### status_code<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `pending-acceptance`, `failed`, `expired`, `provisioning`, `active`, `deleting`, `deleted` or `rejected`.
### cdc_id<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cluster Data Centre

## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_aws_vpc_peer_v2.<resource-name> "<resource-id>"
```
`<resource-id>` is the unique identifier for this resource matching the value of the `id` property defined above.
