provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "dc_test_cluster" {
  cluster_name = "dc_test_cluster"
  node_size = "t3.small-v2"
  data_centre = "US_EAST_1"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "10.0.0.0/16"
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
    version = "3.11.8"
    options = {
      auth_n_authz = true
      use_private_broadcast_rpc_address = false
      client_encryption = false
      lucene_enabled = false
      continuous_backup_enabled = true
    }
  }
}


