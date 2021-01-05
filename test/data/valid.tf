provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
    cluster_name = "test_cluster"
    node_size = "t3.small-v2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 5
        nodes_per_rack = 2
    }

    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.8"
        options = {
            auth_n_authz = true
            use_private_broadcast_rpc_address = true
            lucene_enabled = true
            continuous_backup_enabled = true
        }
    }

    bundle {
        bundle = "SPARK"
        version = "apache-spark:2.3.2"
    }
}
