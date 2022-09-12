---
page_title: "instaclustr_azure_vnet_peer_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_vnet_peer_v2_instance (Data Source)
Definition of an Azure Virtual Network Peering request to allow privately routed connections to a target data centre.
## Example Usage
```
data "instaclustr_azure_vnet_peer_v2_instance" "example" { 
  id = "<id>" // the value of the `id` attribute defined in the root schema below
}
```
## Glossary
The following terms are used to describe attributes in the schema of this data source:
- **_read-only_** - These are attributes that can only be read and not provided as an input to the data source.<br><br>
- **_required_** - These attributes must be provided for the data source's information to be queried.<br><br>
- **_nested block_** - These attributes use the [Terraform block syntax](https://www.terraform.io/language/attr-as-blocks) when defined as an input in the Terraform code. Attributes with the type **_repeatable nested block_** are the same except that the nested block can be defined multiple times with varying nested attributes. When reading nested block attributes, an index must be provided when accessing the contents of the nested block, example - `my_resource.nested_block_attribute[0].nested_attribute`.
# Schema
## Read-only attributes
### peer_subnets<br>
<ins>Type</ins>: list of strings, read-only<br>
<br>The subnets for the peering VPC.
### cdc_id<br>
<ins>Type</ins>: string (uuid), read-only<br>
<br>ID of the Cluster Data Centre.
### status_code<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `GENESIS`, `PROVISIONING`, `FAILED`, `INACTIVE`, `CONNECTED` or `UNKNOWN`.
### id<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.
### name<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the Vpc Peering Connection.
### peer_ad_object_id<br>
<ins>Type</ins>: string, read-only<br>
<br>Id of the Active Directory Object to give peering permissions to, required for cross subscription peering.
### peer_resource_group<br>
<ins>Type</ins>: string, read-only<br>
<br>Resource Group Name of the Virtual Network.
### peer_virtual_network_name<br>
<ins>Type</ins>: string, read-only<br>
<br>The name of the VPC Network you wish to peer to.
### peer_subscription_id<br>
<ins>Type</ins>: string, read-only<br>
<br>Subscription Id of the Virtual Network.
### failure_reason<br>
<ins>Type</ins>: string, read-only<br>
<br>Reason for Peering Connection Failure.
