---
page_title: "instaclustr_vpc_peering Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

### Resource: `instaclustr_vpc_peering`  
A resource for managing VPC peering connections on Instaclustr Managed Platform. This is only avaliable for clusters hosted with the AWS provider.  
  
When creating this resource, the process will wait for target cluster to be in the `PROVISIONED` or `RUNNING` status. The process will time out after 60 seconds of waiting. 

#### Properties
Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr managed cluster|Required
`peer_vpc_id`|The ID of the VPC with which you are creating the VPC peering connection|Required
`peer_account_id`|The account ID of the owner of the accepter VPC|Required
`peer_subnets`|The peer subnets for the VPC|Required
`peer_region`| The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request. | Not Required
`aws_vpc_connection_id`| The ID of the VPC peering connection. | Computed


#### Example
```
resource "instaclustr_vpc_peering" "example_vpc_peering" {
    cluster_id = "${instaclustr_cluster.example.cluster_id}"
    peer_vpc_id = "vpc-123456"
    peer_account_id = "1234567890"
    peer_subnets = ["10.0.0.0/20", "10.0.32.0/20"]
}
```