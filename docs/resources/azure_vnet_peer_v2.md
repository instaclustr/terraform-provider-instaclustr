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
- **_read-only_** - These are attributes that can only be read and not provided as an input to the resource.<br><br>
- **_required_** - These attributes must be provided for the resource to be created.<br><br>
- **_optional_** - These input attributes can be omitted, and doing so may result in a default value being used.<br><br>
- **_immutable_** - These are input attributes that cannot be changed after the resource is created.<br><br>
- **_updatable_** - These input attributes can be updated to a different value if needed, and doing so will trigger an update operation.<br><br>
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
# Schema
## Input attributes - Required
### peer_subnets<br>
<ins>Type</ins>: list of strings, required, immutable<br>
<br>The subnets for the peering VPC.
### cdc_id<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cluster Data Centre.
### peer_resource_group<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Resource Group Name of the Virtual Network.
### peer_virtual_network_name<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The name of the VPC Network you wish to peer to.
### peer_subscription_id<br>
<ins>Type</ins>: string, required, immutable<br>
<br>Subscription Id of the Virtual Network.
## Input attributes - Optional
### peer_ad_object_id<br>
<ins>Type</ins>: string, optional, immutable<br>
<br>Id of the Active Directory Object to give peering permissions to, required for cross subscription peering.
## Read-only attributes
### status_code<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `GENESIS`, `PROVISIONING`, `FAILED`, `INACTIVE`, `CONNECTED` or `UNKNOWN`.
### id<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.
### name<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the Vpc Peering Connection.
### failure_reason<br>
<ins>Type</ins>: string, read-only<br>
<br>Reason for Peering Connection Failure.
## Import
This resource can be imported using the `terraform import` command as follows:
```
terraform import instaclustr_azure_vnet_peer_v2.<resource-name> "<resource-id>"
```
`<resource-id>` is the unique identifier for this resource matching the value of the `id` attribute defined in the root schema above.
