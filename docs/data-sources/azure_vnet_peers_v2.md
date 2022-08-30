---
page_title: "instaclustr_azure_vnet_peers_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_azure_vnet_peers_v2 (Data Source)
A listable data source of all Azure Virtul Network Peering requests in an Instaclustr Account.

## Schema
### account_id<br>
<ins>Type</ins>: string<br>

### peering_requests<br>
<ins>Type</ins>: repeatable nested block, see [peering_requests](#nested--peering_requests) for nested schema<br>

<a id="nested--peering_requests"></a>
## Nested schema for `peering_requests`<br>

### peer_virtual_network_name<br>
<ins>Type</ins>: string<br>
<br>The name of the VPC Network you wish to peer to.
### peer_subnets<br>
<ins>Type</ins>: list of arrays<br>
<br>The subnets for the peering VPC.
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
### cdc_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the Cluster Data Centre.

