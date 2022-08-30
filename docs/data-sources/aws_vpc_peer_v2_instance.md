---
page_title: "instaclustr_aws_vpc_peer_v2_instance Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
---

# instaclustr_aws_vpc_peer_v2_instance (Data Source)

## Schema
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
### status_code<br>
<ins>Type</ins>: string<br>
<br>Status of the VPC Peering Connection. Values can be `pending-acceptance`, `failed`, `expired`, `provisioning`, `active`, `deleting`, `deleted` or `rejected`.
### cdc_id<br>
<ins>Type</ins>: string (uuid)<br>
<br>ID of the Cluster Data Centre

