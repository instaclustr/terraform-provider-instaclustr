// This is part of testing "kafka user" suite, 1 of 3
provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "kafka_cluster" {
  cluster_name = "example_kafka_tf_test"
  node_size = "t3.small-20-gp2"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "KAFKA"
    version = "apache-kafka:2.5.1"
    options = {
      auto_create_topics = true
      client_encryption = false
      dedicated_zookeeper = true
      delete_topics = true
      zookeeper_node_size = "%s"
      zookeeper_node_count = 3
    }
  }
}
