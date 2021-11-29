---
page_title: "instaclustr_kafka_topic Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Data Source `instaclustr_kafka_topic_list`
A read-only data source used to get the list of Kafka topics in a Kafka cluster.


## Properties


`instaclustr_kafka_topic_list`

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr Kafka managed cluster. |Required

## Example
```terraform
data "instaclustr_kafka_topic_list" "kafka_topic_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```
## Returned example in terraform.tfstate
```
"attributes": {
"cluster_id": "cc30db44-ba8a-464a-94b2-8b0cfa70052b",
"id": "cc30db44-ba8a-464a-94b2-8b0cfa70052b-topic-list",
"topics": [
  "__consumer_offsets",
  "instaclustr-sla",
  "test",
  "test2",
  "test3"
]
}
```
