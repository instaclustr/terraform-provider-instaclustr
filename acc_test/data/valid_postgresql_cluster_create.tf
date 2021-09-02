provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "validPostgresql" {
  cluster_name = "testcluster"
  node_size = "postgresql-preview-t3.small-v2-5"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }

  bundle {
    bundle = "POSTGRESQL"
    version = "postgresql:13.4"
    options = {
      postgresql_node_count = 1,
      client_encryption = false
    }
  }
}

