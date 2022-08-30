---
page_title: "instaclustr_aws_vpc_peers_v2 Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_vpc_peers_v2 (Data Source)

## Schema
### account_id<br>
<ins>Type</ins>: string<br>

### peering_requests<br>
<ins>Type</ins>: block list, see [peering_requests](#nested--peering_requests) for nested schema<br>

<a id="nested--peering_requests"></a>
## Nested schema for `peering_requests`<br>

### peer_aws_account_id<br>
<ins>Type</ins>: string<br>
<br>The AWS account ID of the owner of the accepter VPC.
### peer_subnets<br>
<ins>Type</ins>: list of arrays<br>
<br>The subnets for the peering VPC.
### peer_vpc_id<br>
<ins>Type</ins>: string<br>
<br>ID of the VPC with which the peering connection is created.
### id<br>
<ins>Type</ins>: string<br>
<br>ID of the VPC peering connection.
### peer_region<br>
<ins>Type</ins>: string<br>
<br>Region code for the accepter VPC.
### cdc_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the Cluster Data Centre

