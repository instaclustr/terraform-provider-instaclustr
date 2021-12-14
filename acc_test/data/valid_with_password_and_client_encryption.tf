provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "valid_with_password_and_client_encryption" {
  cluster_name = "tf-provider-test-auth-n-ce"
  node_size = "m5l-250-v2"
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
    version = "3.11.8"
    options = {
      auth_n_authz = true
      use_private_broadcast_rpc_address = true
      lucene_enabled = true
      continuous_backup_enabled = true
      client_encryption = true
    }
  }
}

data "instaclustr_cluster_credentials" "cluster_credentials" {
  cluster_id = "${instaclustr_cluster.valid_with_password_and_client_encryption.id}"
}
