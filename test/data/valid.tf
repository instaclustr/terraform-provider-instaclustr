provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "valid" {
    cluster_name = "testcluster"
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
        },
        {
            bundle = "SPARK"
            version = "apache-spark:2.3.2"
        },
        {
            bundle = "ZEPPELIN"
            version = "apache-zeppelin:0.8.0-spark-2.3.2"
        }
    ]
}
