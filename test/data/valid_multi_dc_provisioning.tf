provider "instaclustr" {
  username = "%s"
  api_key = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
  cluster_name = "cassandra-5dc"
  node_size = "t3.small-v2"
  sla_tier = "NON_PRODUCTION"
  cluster_provider = {
    name = "AWS_VPC"
  }
  rack_allocation = {
    number_of_racks = 1
    nodes_per_rack = 2
  }

  data_center {
    data_center = "US_WEST_1"
    network = "192.168.0.0/18"
  }
  data_center {
    data_center = "US_WEST_2"
    network = "10.0.0.0/18"
  }
  data_center {
    name = "backup1"
    data_center = "US_WEST_2"
    network = "10.1.0.0/18"
  }
  data_center {
    name = "backup2"
    data_center = "US_WEST_2"
    network = "10.2.0.0/18"
  }
  data_center {
    name = "backup3"
    data_center = "US_WEST_2"
    network = "10.3.0.0/18"
  }

  bundle {
    bundle = "APACHE_CASSANDRA"
    version = "apache-cassandra-3.11.8.ic2"
    options = {
      auth_n_authz = true
      continuous_backup_enabled = true
    }
  }
}
