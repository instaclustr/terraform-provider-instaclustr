// This is part of testing "kafka acl" suite, 3 of 3
data "instaclustr_kafka_acl_list" "test_acl_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}

