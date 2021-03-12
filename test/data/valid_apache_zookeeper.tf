provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}
resource "instaclustr_cluster" "validApacheZookeeper" {
    cluster_name = "test_zookeeper_cluster"
    node_size = "zookeeper-developer-t3.small-20"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false

    cluster_provider = {
        name = "AWS_VPC"
    }

    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }
    
    bundle {
        bundle = "APACHE_ZOOKEEPER"
        version = "apache-zookeeper:3.5.8"
        options = {
            zookeeper_node_count = 3
        }
    }
}
