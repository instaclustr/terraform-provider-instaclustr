// This is part of testing "kafka acl" suite, 2 of 3 
resource "instaclustr_kafka_acl" "test_acl_2" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  principal = "%s"
  host = "%s"
  resource_type = "%s"
  resource_name = "%s"
  operation = "%s"
  permission_type = "%s"
  pattern_type = "%s"
}

