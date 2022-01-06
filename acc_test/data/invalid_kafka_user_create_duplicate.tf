// This is part of testing "kafka user" suite, 5 of 5 
resource "instaclustr_kafka_user" "kafka_user_charlie_double" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  username = "%s"
  password = "%s"
  initial_permissions = "none"
  override_existing_user = false
}

