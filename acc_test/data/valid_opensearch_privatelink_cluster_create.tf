provider "instaclustr" {
  username     = "%s"
  api_key      = "%s"
  api_hostname = "%s"
}


resource "instaclustr_cluster" "validOpenSearchPrivateLink" {
  cluster_name            = "tf-opensearch-test"
  data_centre             = "US_WEST_2"
  sla_tier                = "NON_PRODUCTION"
  cluster_network         = "192.168.0.0/18"
  private_network_cluster = true
  cluster_provider        = {
    name = "AWS_VPC",
  }
  rack_allocation         = {
    number_of_racks = 3
    nodes_per_rack  = 2
  }
  bundle {
    bundle  = "OPENSEARCH"
    version = "1.3.5"
    options = {
      dedicated_master_nodes  = false
      master_node_size        = "SRH-PRD-m6g.large-120"
      data_node_size          = "SRH-PRD-m6g.large-10"
      security_plugin         = true
      index_management_plugin = true
      knn_plugin              = true
      reporting_plugin        = true
      icu_plugin              = true
      alerting_plugin         = false
      notifications_plugin    = false
      client_encryption       = true
    }
  }

  private_link {
    iam_principal_arns = ["arn:aws:iam::123456789012:root"]
  }
}
