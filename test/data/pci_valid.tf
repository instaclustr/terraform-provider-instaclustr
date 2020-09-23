provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "valid" {
    cluster_name = "testcluster"
    node_size = "m5l-250-v2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = true
    pci_compliant_cluster = true
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }

    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.4"
        options = {
            auth_n_authz = true
            use_private_broadcast_rpc_address = true
            client_encryption = true
        }
    }
}
