provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "validRedis" {
  cluster_name = "testcluster"
  node_size = "t3.small-20-r"
  data_centre = "US_WEST_2"
  sla_tier = "NON_PRODUCTION"
  cluster_network = "192.168.0.0/18"
  cluster_provider = {
    name = "AWS_VPC"
  }

  bundle {
    bundle = "REDIS"
    version = "6.0.4"
    options = {
      master_nodes = 3,
      replica_nodes = 3
    }
  }
}

