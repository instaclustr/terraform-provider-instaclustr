---
page_title: "instaclustr_gcp_vpc_peers_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_gcp_vpc_peers_v2 (Data Source)

## Schema
### account_id<br>
<ins>Type</ins>: string<br>

### peering_requests<br>
<ins>Type</ins>: block list, see [peering_requests](#nested--peering_requests) for nested schema<br>

<a id="nested--peering_requests"></a>
## Nested schema for `peering_requests`<br>

### peer_project_id<br>
<ins>Type</ins>: string<br>
<br>The project ID of the owner of the accepter VPC.
### peer_subnets<br>
<ins>Type</ins>: list of arrays<br>
<br>The subnets for the peering VPC.
### name<br>
<ins>Type</ins>: string<br>
<br>Name of the Peering Connection.
### id<br>
<ins>Type</ins>: string<br>
<br>ID of the VPC peering connection.
### peer_vpc_network_name<br>
<ins>Type</ins>: string<br>
<br>The name of the VPC Network you wish to peer to.
### cdc_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the Cluster Data Centre.

