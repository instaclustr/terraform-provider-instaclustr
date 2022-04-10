// This is part of testing "kafka user" suite, 4 of 5
resource "instaclustr_kafka_user" "kafka_user_charlie_invalid" {
  cluster_id          = "${instaclustr_cluster.kafka_cluster.id}"
  username            = "%s"
  password            = "%s"
  initial_permissions = "none"
  authentication_mechanism = "ExpectedToFail"
}
