---
page_title: "instaclustr_kafka_topic Resource - terraform-provider-instaclustr"
subcategory: ""
description: |-
  
---

# Resource  `instaclustr_kafka_topic` and Data Source `instaclustr_kafka_topic_list`
Resources for managing Kafka topics for a Kafka cluster.
Kafka topic list is a read-only data source used to get the list of kafka topics in a cluster,
while Kafka topic is a resource used to create topics, read topics' replication-factor and partitions, 
read topics' configs, update topics' configs, and delete Kafka topics.

**_Note that we don't allow an update of a topic's replication factor and partitions number._**

## Import
You can import the existing topics of a Kafka cluster created via other ways (e.g. API, Console) to your
terraform resources.

### Import a topic and its replication-factor and partitions
```shell
terraform import instaclustr_kafka_topic.<resource-name> "<cluster-id>&<topic-name>"
```

### Import a topic and its configs
```shell
terraform import instaclustr_kafka_topic.<resource-name> "<cluster-id>&<topic-name>&config"
```
Replace `<resource-name>`, `<cluster-id>`, and `<topic-name>` with the real ones, keep this "&" symbol in the command.

### Example
```shell
terraform import instaclustr_kafka_topic.kafka_topic_test "f4cb7a63-8217-4dc8-ab12-28e54efc00d1&test"
terraform import instaclustr_kafka_topic.kafka_topic_test "f4cb7a63-8217-4dc8-ab12-28e54efc00d1&test&config"
```

After importing, you should see the imported topic in ***terraform.tfstate***.

## Schema
Property | Description | Default
---------|-------------|--------
cluster_id | The ID of an existing Instaclustr Kafka managed cluster. | Required
topic | The topic name for the Kafka topic. | Required
replication_factor | The replication factor of this topic. | Required
partitions | The partitions number of this topic. | Required
config | The list of this topic's configs, see below for its properties. | Optional, but required if you want to update the topic's configs

### config
This contains all the possible configs that you can change for a Kafka topic. We use Kafka AdminClient 2.3.1 internally, please see https://kafka.apache.org/23/documentation.html#topicconfigs
to check every config's information. 

Property | Data Type | Default
---------|-------------|--------
compression_type | string | None
leader_replication_throttled_replicas | string | ""
min_insync_replicas | int | None
message_downconversion_enable | boolean | true
segment_jitter_ms | int | 0
cleanup_policy | string | None
flush_ms | string<sup>* | None
follower_replication_throttled_replicas | string | None
retention_ms | int | None
segment_bytes | int | None
flush_messages | string<sup>* | None
message_format_version| string | None
file_delete_delay_ms | int | None
max_compaction_lag_ms | string<sup>* | None
max_message_bytes | int | None
min_compaction_lag_ms | int | 0
message_timestamp_type | string | None
preallocate | boolean | false
index_interval_bytes | int | None
min_cleanable_dirty_ratio | float | None
unclean_leader_election_enable | boolean | false
delete_retention_ms | int | None
retention_bytes | int | None
segment_ms | int | None
message_timestamp_difference_max_ms | string<sup>* | None
segment_index_bytes | int | None

<sup>* Indeed an int value, using string because Terraform has difficulties parsing big int, it will lose precision.

## Usage
There are two use cases for updating a Kafka topic's configs.
1. Providing all the possible configs in resource, and modify the ones you want to update. Then keeping all the configs in the resource.
2. Providing some configs (not all) you want to update, and then keep only these configs in the resource.

Terraform always detects that there is an update in case 2 if you run the plan again. Although Terraform detects an update, we make sure
that there will not be a real update request if the configs in the resources are not changed.

### Case 1 example
```terraform
resource "instaclustr_kafka_topic" "kafka_topic_test2" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  topic = "test"
  replication_factor = 3
  partitions = 3
  config {
    cleanup_policy = "delete"
    compression_type = "producer"
    delete_retention_ms = 86400000
    file_delete_delay_ms = 60000
    flush_messages = "9223372036854775807"
    flush_ms = "9223372036854775807"
    follower_replication_throttled_replicas = ""
    index_interval_bytes = 4096
    leader_replication_throttled_replicas = ""
    max_compaction_lag_ms = "9223372036854775807"
    max_message_bytes = 1048588
    message_downconversion_enable = true
    message_format_version = "2.3-IV1" // This is changed from "3.0-IV1" -> "2.3-IV1"
    message_timestamp_difference_max_ms = "9223372036854775807"
    message_timestamp_type = "CreateTime"
    min_cleanable_dirty_ratio = 0.5
    min_compaction_lag_ms = 0
    min_insync_replicas = 2 // This is changed from 1 -> 2
    preallocate = false
    retention_bytes = -1
    retention_ms = 604800000
    segment_bytes = 1073741824
    segment_index_bytes = 10485760
    segment_jitter_ms = 0
    segment_ms = 604800000
    unclean_leader_election_enable = true // This is changed from false -> true
  }
}

data "instaclustr_kafka_topic_list" "kafka_topic_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```

### Case 2 example
```terraform
resource "instaclustr_kafka_topic" "kafka_topic_test2" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  topic = "test"
  replication_factor = 3
  partitions = 3
  config {
    message_format_version = "2.3-IV1" // This is changed from "3.0-IV1" -> "2.3-IV1"
    min_insync_replicas = 2 // This is changed from 1 -> 2
    preallocate = true // This is changed from false -> true
    unclean_leader_election_enable = true // This is changed from false -> true
  }
}

data "instaclustr_kafka_topic_list" "kafka_topic_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}
```