---
page_title: "instaclustr_vpc_peering Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

### Resource: `instaclustr_vpc_peering`  
A resource for managing VPC peering connections on Instaclustr Managed Platform. This is avaliable for clusters hosted with AWS and GCP.
  
When creating this resource, the process will wait for target cluster to be in the `PROVISIONED` or `RUNNING` status. The process will time out after 60 seconds of waiting. 

#### Properties for AWS
Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`peer_vpc_id`|The ID of the VPC with which you are creating the VPC peering connection|Required
`peer_account_id`|The account ID of the owner of the accepter VPC|Required
`peer_subnets`|A set of one or more peer subnets (in IPv4 CIDR format) for the VPC|Required
`peer_region`| The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request. | Not Required
`aws_vpc_connection_id`| The ID of the VPC peering connection. | Computed

#### Properties for GCP
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`peer_vpc_network_name`|The ID of the VPC with which you are creating the VPC peering connection|Required
`peer_project_id`|The project ID of the owner of the accepter VPC|Required
`peer_subnets`|A set of one or more peer subnets (in IPv4 CIDR format) for the VPC|Required

#### Properties for Azure
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`peer_subscription_id`|The Suscription ID of the VPC with which you are creating the VPC peering connection|Required
`peer_resource_group`|The Resource Group name with which you are creating the VPC peering connection|Required
`peer_vpc_net`|The name of the VPC with which you are creating the VPC peering connection|Required


#### Example
```
resource "instaclustr_vpc_peering" "example_vpc_peering" {
    cluster_id = "${instaclustr_cluster.example.cluster_id}"
    peer_vpc_id = "vpc-123456"
    peer_account_id = "1234567890"
    peer_subnets = toset(["10.0.0.0/20", "10.0.32.0/20"])
}
```

#### Legacy format (Only for AWS)
The legacy VPC peering connection type only supports a single IPv4 CIDR for the peer subnet field. To maintain backwards compatibility this resource definition is still available; however we recommend against using it for new plans.
Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`peer_vpc_id`|The ID of the VPC with which you are creating the VPC peering connection|Required
`peer_account_id`|The account ID of the owner of the accepter VPC|Required
`peer_subnet`|The peer subnet (in IPv4 CIDR format) for the VPC|Required
`peer_region`| The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request. | Not Required
`aws_vpc_connection_id`| The ID of the VPC peering connection. | Computed

#### AWS Example
```
resource "instaclustr_vpc_peering" "example_vpc_peering" {
    cluster_id = "${instaclustr_cluster.example.cluster_id}"
    peer_vpc_id = "vpc-123456"
    peer_account_id = "1234567890"
    peer_subnet = "10.0.0.0/20"
}
```

#### GCP Example
```
resource "instaclustr_vpc_peering_gcp" "gcp_example" {
  peer_vpc_network_name = "my-vpc1"
  peer_project_id = "project-id"
  peer_subnets = toset(["10.10.0.0/16", "10.11.0.0/16"])
  cluster_id = "${instaclustr_cluster.gcp_example.id}"
}

```
#### Azure Example
```
resource "instaclustr_vpc_peering_azure" "example_vpc_peering" {
  cluster_id = "c820f98f-06b9-491b-8504-58694ea7e57f"
  peer_subscription_id="7a07f268-eb64-45df-b63e-b5595e713287"
  peer_resource_group="instaclustrtest"
  peer_vpc_net="nana"
  peer_subnets = toset(["10.8.0.0/16", "10.11.0.0/16"])
  
}

```