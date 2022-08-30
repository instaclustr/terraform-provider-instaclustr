---
page_title: "instaclustr_gcp_vpc_peer_v2 Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_gcp_vpc_peer_v2 (Resource)

## Schema
### peer_project_id<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The project ID of the owner of the accepter VPC.
### peer_subnets<br>
<ins>Type</ins>: list of arrays, required, immutable<br>
<br>The subnets for the peering VPC.
### failure_reason<br>
<ins>Type</ins>: string, read-only<br>
<br>Reason for Peering Connection Failure.
### name<br>
<ins>Type</ins>: string, read-only<br>
<br>Name of the Peering Connection.
### id<br>
<ins>Type</ins>: string, read-only<br>
<br>ID of the VPC peering connection.
### peer_vpc_network_name<br>
<ins>Type</ins>: string, required, immutable<br>
<br>The name of the VPC Network you wish to peer to.
### status_code<br>
<ins>Type</ins>: string, read-only<br>
<br>Status of the VPC Peering Connection. Values can be `GENESIS`, `PROVISIONING`, `FAILED`, `INACTIVE`, `ACTIVE` or `UNKNOWN`.
### cdc_id<br>
<ins>Type</ins>: string (uuid), required, immutable<br>
<br>ID of the Cluster Data Centre.

