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
    version = "opensearch:1.0.0.ic2"
    options = {
      dedicated_master_nodes          = false,
      master_node_size                = "t3.small-v2",
      opensearch_dashboards_node_size = "t3.small-v2",
      security_plugin                 = true,
      client_encryption               = true
    }
  }
}