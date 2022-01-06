// This is part of testing "kafka topic" suite, 2 of 4
data "instaclustr_kafka_topic_list" "kafka_topic_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}