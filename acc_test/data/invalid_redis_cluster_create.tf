provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "invalidRedis" {
  cluster_name = "tf-redis-test"
  node_size = "t3.small-20-r"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = true
  pci_compliant_cluster   = false
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    nodes_per_rack = 1
    number_of_racks = 4
  }

  bundle {
    bundle = "REDIS"
    version = "redis:6.0.9"
    options = {
      master_nodes = 3,
      replica_nodes = 3,
      password_auth = false,
      client_encryption = false
    }
  }

  wait_for_state = "RUNNING"
}

