provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "vpc_cluster" {
    cluster_name = "vpc_cluster"
    node_size = "t3.small-v2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC",
        account_name = "%s"
        custom_virtual_network_id = "%s"
    }
    rack_allocation = {
        number_of_racks = 2
        nodes_per_rack = 1
    }

    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "apache-cassandra-3.11.8.ic2"
        options = {
            auth_n_authz = false
            client_encryption = false
            continuous_backup_enabled = false
            lucene_enabled = false
            use_private_broadcast_rpc_address = false
        }
    }
}
