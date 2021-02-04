provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}
resource "instaclustr_cluster" "valid" {
    cluster_name = "test_cluster"
    node_size = "t3.small-20-gp2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 5
        nodes_per_rack = 2
    }

    bundle {
        bundle = "KAFKA"
        version = "apache-kafka:2.5.1"
        options = {
            client_encryption = false
            auto_create_topics = true
            delete_topics = true
            dedicated_zookeeper = false
        }
    }
}
