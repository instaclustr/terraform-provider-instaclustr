provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}

resource "instaclustr_cluster" "resizable_cluster" {
  cluster_name            = "tf-resizable-test"
  data_centre             = "US_EAST_1"
  sla_tier                = "NON_PRODUCTION"
  cluster_network         = "192.168.0.0/18"
  private_network_cluster = false
  wait_for_state          = "RUNNING"
  cluster_provider        = {
    name = "AWS_VPC"
  }
  rack_allocation         = {
    number_of_racks = 3
    nodes_per_rack  = 1
  }
  bundle {
    bundle  = "OPENSEARCH"
    version = "1.3.4"
    options = {
      dedicated_master_nodes          = false,
      master_node_size                = "SRH-DEV-t4g.small-5",
      opensearch_dashboards_node_size = "SRH-DEV-t4g.small-5",
      security_plugin                 = true,
      client_encryption               = true
    }
  }
}
