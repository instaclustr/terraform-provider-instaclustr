provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
  cluster_name = "testcluster"
  node_size = "t3.small-v2"
  sla_tier = "NON_PRODUCTION"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 2
    nodes_per_rack = 1
  }


  data_centres {
    data_centre_region = "US_WEST_1"
    network = "10.1.0.0/18"
  }


  data_centres {
        data_centre_region = "US_WEST_2"
        network = "192.168.0.0/18"
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "apache-cassandra-3.11.8.ic2"
    options = {
      auth_n_authz = true
      continuous_backup_enabled = true
      client_encryption = false
      lucene_enabled = false
      use_private_broadcast_rpc_address = false
    }
  }
}
