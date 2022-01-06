// This is part of testing "kafka topic" suite, 4 of 4
resource "instaclustr_kafka_topic" "kafka_topic_test" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  topic = "%s"
  replication_factor = 3
  partitions = 3
}

resource "instaclustr_kafka_topic" "kafka_topic_test2" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  topic = "%s"
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