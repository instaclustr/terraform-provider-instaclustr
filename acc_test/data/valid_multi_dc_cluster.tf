provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "dc_test_cluster" {
  cluster_name            = "dc_test_cluster"
  cluster_network         = "10.0.0.0/16"
  pci_compliant_cluster   = false
  private_network_cluster = false
  sla_tier                = "NON_PRODUCTION"

  data_centres {
    data_centre     = "US_EAST_1"
    name            = "US_EAST_1"
    network         = "10.0.0.0/16"
    node_size       = "t3.small-v2"
    provider        = {
      "name" = "AWS_VPC"
    }
    rack_allocation = {
      "nodes_per_rack"  = "1"
      "number_of_racks" = "3"
    }

    bundles {
      bundle  = "APACHE_CASSANDRA"
      options = {
        "auth_n_authz"                      = "true"
        "client_encryption"                 = "false"
        "continuous_backup_enabled"         = "true"
        "lucene_enabled"                    = "false"
        "use_private_broadcast_rpc_address" = "false"
      }
      version = "apache-cassandra-3.11.8.ic2"
    }
  }

  data_centres {
    data_centre     = "US_WEST_1"
    name            = "DC2"
    network         = "10.1.0.0/16"
    node_size       = "t3.small-v2"
    provider        = {
      "name" = "AWS_VPC"
    }
    rack_allocation = {
      "nodes_per_rack"  = "1"
      "number_of_racks" = "3"
    }

    bundles {
      bundle  = "APACHE_CASSANDRA"
      options = {
        "auth_n_authz"                      = "true"
        "client_encryption"                 = "false"
        "continuous_backup_enabled"         = "true"
        "lucene_enabled"                    = "false"
        "use_private_broadcast_rpc_address" = "false"
      }
      version = "apache-cassandra-3.11.8.ic2"
    }
  }
}