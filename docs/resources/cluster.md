---
page_title: "instaclustr_cluster Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource `instaclustr_cluster`





## Schema

### Required

- **bundle** (Block List, Min: 1) (see [below for nested schema](#nestedblock--bundle))
- **cluster_name** (String)
- **cluster_provider** (Map of String)
- **data_centre** (String)
- **node_size** (String)

### Optional

- **bundles** (List of Map of String)
- **cluster_network** (String)
- **id** (String) The ID of this resource.
- **kafka_rest_proxy_user_password** (String, Sensitive)
- **kafka_schema_registry_user_password** (String, Sensitive)
- **pci_compliant_cluster** (Boolean)
- **private_contact_point** (String)
- **private_network_cluster** (Boolean)
- **public_contact_point** (String)
- **rack_allocation** (Map of String)
- **sla_tier** (String)
- **tags** (Map of String)
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **wait_for_state** (String)

### Read-only

- **cluster_id** (String)

<a id="nestedblock--bundle"></a>
### Nested Schema for `bundle`

Required:

- **bundle** (String)
- **version** (String)

Optional:

- **options** (Map of String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String)


