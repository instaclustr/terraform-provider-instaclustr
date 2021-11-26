---
page_title: "insbtaclustr_kafka_acl_list Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Data Source `instaclustr_kafka_acl_list`                             
Resources for managing Kafka ACL for a Kafka cluster. 
Kafka ACL list is a read-only data source used to get the list of kafka ACL in a cluster, 

## Properties

`instaclustr_kafka_acl_list`

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr Kafka managed cluster. |Required

## Example

```
data "instaclustr_kafka_acl_list" "kafka_acl_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```
