---
page_title: "instaclustr_kafka_topic Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Data Source `instaclustr_kafka_topic_list`
Resources for managing Kafka topics for a Kafka cluster.
Kafka topic list is a read-only data source used to get the list of kafka topics in a cluster,


#### Properties


`instaclustr_kafka_topic_list`

Property | Description | Default
---------|-------------|--------
`cluster_id`|The ID of an existing Instaclustr Kafka managed cluster. |Required

#### Example
```terraform
data "instaclustr_kafka_topic_list" "kafka_topic_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```

