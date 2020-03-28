provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "vpc_cluster" {
    cluster_name = "vpc_cluster"
    node_size = "t3.small"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC",
        account_name = "%s"
        custom_virtual_network_id = "vpc-asdasdasd"
    }
    rack_allocation = {
        number_of_racks = 2
        nodes_per_rack = 1
    }

    bundle {
        bundle = "APACHE_CASSANDRA"
        version = "3.11.4"
    }
}
