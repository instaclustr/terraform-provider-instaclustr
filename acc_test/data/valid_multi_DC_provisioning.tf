provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
  cluster_name = "testcluster"
  sla_tier     = "NON_PRODUCTION"

  data_centres {
    name        = "DC1"
    data_centre = "US_WEST_1"
    network     = "10.0.0.0/18"
    node_size    = "m5l-250-v2"
    rack_allocation = {
      number_of_racks = 2
      nodes_per_rack  = 1
    }
    provider = {
      name = "AWS_VPC"
    }

    bundles {
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

    bundles {
      bundle = "SPARK"
      version = "apache-spark:2.3.2.ic3"
    }
  }

  data_centres {
    name        = "DC2"
    data_centre = "US_WEST_1"
    network     = "10.1.0.0/18"
    node_size    = "m5l-250-v2"
    rack_allocation = {
      number_of_racks = 2
      nodes_per_rack  = 1
    }
    provider = {
      name = "AWS_VPC"
    }
    bundles {
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

  data_centres {
    data_centre = "US_WEST_2"
    network     = "192.168.0.0/18"
    node_size    = "m5l-250-v2"
    rack_allocation = {
      number_of_racks = 3
      nodes_per_rack  = 1
    }
    provider = {
      name = "AWS_VPC"
    }
    bundles {
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
    bundles {
      bundle = "SPARK"
      version = "apache-spark:2.3.2.ic3"
    }
  }
}
