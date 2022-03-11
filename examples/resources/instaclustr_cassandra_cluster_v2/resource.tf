resource "instaclustr_cassandra_cluster_v2" "test_cluster" {
  name                    = "MyTestCluster"
  cassandra_version       = "4.0.1"
  lucene_enabled          = true
  password_and_user_auth  = true
  pci_compliance_mode     = false
  private_network_cluster = false
  sla_tier                = "PRODUCTION"
  spark_version           = "3.0.1"

  data_centre {
    name                               = "MyTestDataCentre"
    client_to_cluster_encryption       = true
    cloud_provider                     = "AWS_VPC"
    continuous_backup                  = true
    node_size                          = "m5l-250-v2"
    number_of_nodes                    = 2
    replication_factor                 = 2
    private_ip_broadcast_for_discovery = true
    network                            = "10.0.0.0/16"
    region                             = "US_EAST_1"
  }

  two_factor_delete {
    confirmation_email = "test@email.com"
  }
}
