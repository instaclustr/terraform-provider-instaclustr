provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}


resource "instaclustr_cluster" "invalidElasticsearch" {
    cluster_name = "tf-resizable-elasticsearch-test"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC",
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    bundle {
        bundle = "ELASTICSEARCH"
        version = "1.11.0"
        options = {
            dedicated_master_nodes = false,
            master_node_size = "m5l-250-v2",
            kibana_node_size = "m5l-400-v2",
            data_node_size = "t3.small-v2",
        }
    }
}