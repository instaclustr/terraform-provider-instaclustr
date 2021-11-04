provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}


resource "instaclustr_cluster" "validElasticsearch" {
    cluster_name = "tf-elasticsearch-test"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    cluster_provider = {
        name = "AWS_VPC",
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 2
    }
    bundle {
        bundle = "ELASTICSEARCH"
        version = "opendistro-for-elasticsearch:1.11.0.ic1"
        options = {
            dedicated_master_nodes = false,
            master_node_size = "t3.small-v2",
            data_node_size = "t3.small-v2",
            security_plugin=false,
            client_encryption=false
        }
    }
}