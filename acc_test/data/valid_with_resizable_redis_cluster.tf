provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "resizable_cluster" {
  cluster_name            = "tf-resizable-test"
  node_size               = "t3.small-20-r"
  data_centre             = "US_EAST_1"
  sla_tier                = "NON_PRODUCTION"
  cluster_network         = "192.168.0.0/18"
  private_network_cluster = false
  pci_compliant_cluster   = false
  wait_for_state          = "RUNNING"
  cluster_provider = {
    name = "AWS_VPC"
  }
  bundle {
    bundle  = "REDIS"
    version = "redis:6.0.9"
    options = {
      master_nodes      = 3,
      replica_nodes     = 3,
      password_auth     = false,
      client_encryption = false,
    }
  }
}