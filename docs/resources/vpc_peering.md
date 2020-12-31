---
page_title: "instaclustr_vpc_peering Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource `instaclustr_vpc_peering`





## Schema

### Required

- **cluster_id** (String)
- **peer_account_id** (String)
- **peer_subnet** (String)
- **peer_vpc_id** (String)

### Optional

- **id** (String) The ID of this resource.
- **peer_region** (String)

### Read-only

- **aws_vpc_connection_id** (String)
- **cdc_id** (String)
- **vpc_peering_id** (String)


