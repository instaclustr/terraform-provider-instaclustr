---
page_title: "instaclustr_clusters Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

### Data Source `instaclustr_clusters`
A read-only data source used to get the list of active clusters under the account that generated the configured API key.

## Attributes Reference

Attribute | Description
---------|-------------
`cluster_id`|The ID of an existing Instaclustr cluster.
`cluster_name`|Cluster name.
`node_count`|Number of nodes in the cluster.
`running_node_count`|Number of nodes in the cluster in the `RUNNING` state.
`derived_status`|Current Status of the cluster. Can be `GENESIS`, `PROVISIONING`, `PROVISIONED`, `JOINING`, `RESTORING` `DEFERRED`, `RUNNING`, `DELETING`, `DELETED`, `DEGRADED` or `FAILED`.
`sla_tier`|SLA tier of the cluster. Can be `PRODUCTION` or `NON_PRODUCTION`.
`pci_compliance`|Cluster PCI compliance. Can be `ENABLED`, `DISABLED` or `PENDING`.
