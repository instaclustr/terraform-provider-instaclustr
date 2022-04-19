provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "validRedis" {
  cluster_name  = "tf-redis-test"
  node_size = "t3.small-20-r"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  pci_compliant_cluster = false
  cluster_provider = {
    name = "AWS_VPC"
  }

  bundle {
    bundle = "REDIS"
    version = "6.2.6"
    options = {
      master_nodes = 3,
      replica_nodes = 3,
      password_auth = false,
      client_encryption = false,
    }
  }

  wait_for_state = "RUNNING"
}

