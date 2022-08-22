provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}


resource "instaclustr_cluster" "invalidOpenSearch" {
  cluster_name            = "tf-resizable-opensearch-test"
  data_centre             = "US_WEST_2"
  sla_tier                = "NON_PRODUCTION"
  cluster_network         = "192.168.0.0/18"
  private_network_cluster = false
  cluster_provider        = {
    name = "AWS_VPC",
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
      master_node_size                = "SRH-DM-m5d.large",
      opensearch_dashboards_node_size = "SRH-PRD-m6g.large-400",
      data_node_size                  = "SRH-DEV-t4g.small-5",
      security_plugin                 = true,
      client_encryption               = true
    }
  }
}
