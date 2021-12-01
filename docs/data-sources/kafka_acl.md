---
page_title: "insbtaclustr_kafka_acl_list Data Source - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Data Source `instaclustr_kafka_acl_list`                             
A data source to read the list of the Kafka ACL configured in a Kafka cluster, 

## Properties

`instaclustr_kafka_acl_list`

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr Kafka managed cluster. |Required

## Attributes Reference
Attribute | Description
----------|------------
`acls`    | List of ACLs in the cluster. Each entry will conform to the structure of something like `(principal=User:*, host=*, resourceType=TOPIC, resourceName=*, operation=ALL, permissionType=ALLOW, patternType=LITERAL)`.

## Example

```
data "instaclustr_kafka_acl_list" "kafka_acl_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```
