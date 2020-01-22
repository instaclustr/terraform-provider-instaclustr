provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "resizable_cluster" {
    cluster_name = "tf-resizable-test"
    node_size = "resizeable-small(r5-l)"
    data_centre = "US_EAST_1"
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
            version = "apache-cassandra-3.0.18"
        }
    ]
}
