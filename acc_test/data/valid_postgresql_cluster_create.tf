provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "validPostgresql" {
  cluster_name = "testcluster"
  node_size = "PGS-DEV-t3.small-5"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    nodes_per_rack = 1
    number_of_racks = 1
  }

  bundle {
    bundle = "POSTGRESQL"
    version = "14.1"
    options = {
      postgresql_node_count = 2,
      client_encryption = false,
      replication_mode = "SYNCHRONOUS",
      synchronous_mode_strict = true
    }
  }
}

