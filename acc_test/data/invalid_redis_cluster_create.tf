provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "validRedis" {
  cluster_name            = "tf-redis-test"
  node_size               = "r5.large-100-r"
  data_centre             = "US_WEST_2"
  sla_tier                = "PRODUCTION"
  cluster_network         = "192.168.0.0/18"
  private_network_cluster = false
  pci_compliant_cluster   = false
  cluster_provider = {
    name = "AWS_VPC"
  }

  rack_allocation = {
    number_of_racks = 5
    nodes_per_rack  = 2
  }

  bundle {
    bundle  = "REDIS"
    version = "6.0.9"
    options = {
      master_nodes      = 3,
      replica_nodes     = 3,
      password_auth     = false,
      client_encryption = false,
    }
  }

  wait_for_state = "RUNNING"
}
