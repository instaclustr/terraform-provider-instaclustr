// This is part of testing "kafka user" suite, 2 of 5
resource "instaclustr_kafka_user" "kafka_user_charlie" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  username = "%s"
  password = "%s"
  initial_permissions = "none"
  override_existing_user = false
}

resource "instaclustr_kafka_user" "kafka_user_charlie_scram-sha-512" {
  cluster_id          = "${instaclustr_cluster.kafka_cluster.id}"
  username            = "%s"
  password            = "%s"
  initial_permissions = "none"
  authentication_mechanism = "SCRAM-SHA-512"
  override_existing_user = false
}

