provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "resizable_pci_cluster" {
    cluster_name = "tf-resizable-test"
    node_size = "resizeable-small(r5-l)-v2"
    data_centre = "US_EAST_1"
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
