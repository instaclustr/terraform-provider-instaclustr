// This is part of testing "kafka user" suite, 3 of 5
provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "kafka_cluster" {
  cluster_name = "example_kafka_tf_test"
  node_size = "%s"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  wait_for_state = "RUNNING"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "KAFKA"
    version = "%s"
    options = {
      auto_create_topics = true
      client_encryption = false
      dedicated_zookeeper = true
      delete_topics = true
      number_partitions = 3
      zookeeper_node_size = "%s"
      zookeeper_node_count = 3
    }
  }
} 

resource "instaclustr_kafka_user" "kafka_user_charlie" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
  username = "%s"
  password = "%s"
  initial_permissions = "none"
}

resource "instaclustr_kafka_user" "kafka_user_charlie_scram-sha-512" {
  cluster_id          = "${instaclustr_cluster.kafka_cluster.id}"
  username            = "%s"
  password            = "%s"
  initial_permissions = "none"
  authentication_mechanism = "SCRAM-SHA-512"
}

data "instaclustr_kafka_user_list" "kafka_user_list" {
  cluster_id = "${instaclustr_cluster.kafka_cluster.id}"
}

