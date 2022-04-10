// This is part of testing "kafka user" suite, 3 of 5
data "instaclustr_kafka_user_list" "kafka_user_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}

