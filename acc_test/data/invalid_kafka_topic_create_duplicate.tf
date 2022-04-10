// This is part of testing "kafka topic" suite, 3 of 4
resource "instaclustr_kafka_topic" "kafka_topic_test3" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  topic = "%s"
  replication_factor = 3
  partitions = 3
}