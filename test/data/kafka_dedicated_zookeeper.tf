provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "kafka_dedicated_zookeeper" {
  cluster_name = "testcluster_dedicatedzookeeper"
  node_size = "t3.small-20-gp2"
  data_centre = "US_EAST_1"
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
}
