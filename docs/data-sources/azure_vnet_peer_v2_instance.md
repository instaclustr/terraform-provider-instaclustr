---
page_title: "instaclustr_azure_vnet_peer_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_vnet_peer_v2_instance (Data Source)
Definition of an Azure Virtual Network Peering request to allow privately routed connections to a target data centre.

## Schema
### peer_virtual_network_name<br>
<ins>Type</ins>: string<br>
<br>The name of the VPC Network you wish to peer to.
### peer_subnets<br>
<ins>Type</ins>: list of arrays<br>
<br>The subnets for the peering VPC.
### failure_reason<br>
<ins>Type</ins>: string<br>
<br>Reason for Peering Connection Failure.
### name<br>
<ins>Type</ins>: string<br>
<br>Name of the Vpc Peering Connection.
### peer_ad_object_id<br>
<ins>Type</ins>: string<br>
<br>Id of the Active Directory Object to give peering permissions to, required for cross subscription peering.
### id<br>
<ins>Type</ins>: string<br>
<br>ID of the VPC peering connection.
### peer_resource_group<br>
<ins>Type</ins>: string<br>
<br>Resource Group Name of the Virtual Network.
### peer_subscription_id<br>
<ins>Type</ins>: string<br>
<br>Subscription Id of the Virtual Network.
### status_code<br>
<ins>Type</ins>: string<br>
<br>Status of the VPC Peering Connection. Values can be `GENESIS`, `PROVISIONING`, `FAILED`, `INACTIVE`, `CONNECTED` or `UNKNOWN`.
### cdc_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the Cluster Data Centre.

