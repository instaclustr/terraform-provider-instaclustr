provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
    cluster_name = "test_cluster"
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
        number_of_racks = 5
        nodes_per_rack = 2
    }

    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "4.0.1"
        options = {
            auth_n_authz = true
            use_private_broadcast_rpc_address = false
            client_encryption = false
            lucene_enabled = false
            continuous_backup_enabled = true
        }
    }
}

resource "instaclustr_cluster" "gcp_valid" {
  cluster_name = "testclustergcp"
  node_size = "n1-standard-2"
  data_centre = "us-east1"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider = {
    name = "GCP"
  }
  rack_allocation = {
    number_of_racks = 3
    nodes_per_rack = 1
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "4.0.1"
    options = {
      auth_n_authz = true
    }
  }
}
