provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "invalidPostgresqlWithPgBouncer" {
  cluster_name    = "testcluster"
  node_size       = "PGS-DEV-t3.small-5"
  data_centre     = "US_WEST_2"
  sla_tier        = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    nodes_per_rack  = 1
    number_of_racks = 1
  }

  bundle {
    bundle  = "POSTGRESQL"
    version = "14.6"
    options = {
      postgresql_node_count = 1,
      client_encryption     = false,
      replication_mode      = "SYNCHRONOUS",
    }
  }
  //Will fail due to invalid pool_mode
  bundle {
    bundle  = "PGBOUNCER"
    version = "1.17.0"
    options = {
      pool_mode = "foobar"
    }
  }
}

