provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "kafka_rest_proxy_cluster" {
  cluster_name = "example_kafka_rest_proxy_tf_test"
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
    version = "2.5.1"
    options = {
      dedicated_zookeeper = true
      zookeeper_node_size = "%s"
      zookeeper_node_count = 3
    }
  }

  bundle {
    bundle = "KAFKA_REST_PROXY"
    version = "5.0.0"
  }
}

resource "instaclustr_bundle_user" "kafka_rest_proxy_user_update" {
  cluster_id = "${instaclustr_cluster.kafka_rest_proxy_cluster.id}"
  username = "ickafkarest"
  password = "%s"
  initial_permissions = "none"
  bundle_name = "kafka_rest_proxy"
}
