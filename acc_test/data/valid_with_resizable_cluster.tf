provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "resizable_cluster" {
    cluster_name = "tf-resizable-test"
    node_size = "t3.medium-v2"
    data_centre = "US_EAST_1"
    data_centre_custom_name = "AWS_VPC_US_EAST_1_name"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    wait_for_state = "RUNNING"
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.8"
        options = {
            auth_n_authz = false
            client_encryption = false
            continuous_backup_enabled = false
            lucene_enabled = false
            use_private_broadcast_rpc_address = false
        }
    }
}
