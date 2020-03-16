provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "valid_with_vpc_peering" {
    cluster_name = "tf_provider_vpc_peering_test"
    node_size = "t3.small"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundles = [
        {
            bundle = "APACHE_CASSANDRA"
            version = "3.11.4"
        }
    ]
}

resource "instaclustr_vpc_peering" "valid_with_vpc_peering" {
    cluster_id = "${instaclustr_cluster.valid_with_vpc_peering.cluster_id}"
    peer_vpc_id = "vpc-12345678"
    peer_account_id = "494111121110"
    peer_subnet = "10.128.176.0/20"
}