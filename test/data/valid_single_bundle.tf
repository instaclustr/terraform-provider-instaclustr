provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
  cluster_name = "testcluster"
  node_size = "t3.small"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  pci_compliant_cluster = false
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "4.0"
    options = {
      auth_n_authz = true
      use_private_broadcast_rpc_address = true
      lucene_enabled = true
      continuous_backup_enabled = true
    }
  }
}
