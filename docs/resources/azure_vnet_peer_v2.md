---
page_title: "instaclustr_azure_vnet_peer_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_vnet_peer_v2 (Resource)
Definition of an Azure Virtual Network Peering request to allow privately routed connections to a target data centre.
## Example Usage
```
resource "instaclustr_azure_vnet_peer_v2" "example" {
  peer_virtual_network_name = "network-aabb-1122"
  peer_subnets = [ "10.1.0.0/16", "10.2.0.0/16" ]
  peer_ad_object_id = "00000000-0000-0000-0000-000000000000"
  peer_resource_group = "example-resource-group-123"
  peer_subscription_id = "00000000-0000-0000-0000-000000000000"
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
<ins>Type</ins>: list of strings, required, immutable<br>
<br>The subnets for the peering VPC.<br><br>
*___cdc_id___*<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cluster Data Centre.<br><br>
*___peer_resource_group___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Resource Group Name of the Virtual Network.<br><br>
*___peer_virtual_network_name___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The name of the VPC Network you wish to peer to.<br><br>
*___peer_subscription_id___*<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Subscription ID of the Virtual Network.<br><br>
### Input attributes - Optional
*___peer_ad_object_id___*<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>ID of the Active Directory Object to give peering permissions to, required for cross subscription peering.<br><br>
### Read-only attributes
*___status_code___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `GENESIS`, `PROVISIONING`, `FAILED`, `INACTIVE`, `CONNECTED` or `UNKNOWN`.<br><br>
*___id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.<br><br>
*___name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the Vpc Peering Connection.<br><br>
*___data_centre_resource_group___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Resource Group Name of the Data Centre Virtual Network.<br><br>
*___data_centre_virtual_network_name___*<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the Data Centre Virtual Network.<br><br>
*___failure_reason___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Reason for Peering Connection Failure.<br><br>
*___data_centre_subscription_id___*<br>
<ins>Type</ins>: string, read-only<br>
<br>Subscription ID of the Data Centre Virtual Network.<br><br>
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_azure_vnet_peer_v2.[resource-name] "[resource-id]"
```
`[resource-id]` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
