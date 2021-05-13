provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}
resource "instaclustr_cluster" "resizable_cluster" {
    cluster_name = "tf-resizable-test"
    node_size = "r5.large-800-gp2"
    data_centre = "US_WEST_2"
    sla_tier = "NON_PRODUCTION"
    cluster_network = "192.168.0.0/18"
    private_network_cluster = false
    pci_compliant_cluster = false
    wait_for_state = "RUNNING"
    cluster_provider = {
        name = "AWS_VPC"
    }
    rack_allocation = {
        number_of_racks = 3
        nodes_per_rack = 1
    }

    bundle {
        bundle = "KAFKA"
        version = "apache-kafka:2.5.1"
        options = {
            client_encryption = false
            auto_create_topics = true
            delete_topics = true
            number_partitions = 10
            zookeeper_node_count = 3
        }
    }

    bundle {
        bundle = "KAFKA_SCHEMA_REGISTRY"
        version = "kafka-schema-registry:5.0.0"
    }
    bundle {
        bundle = "KAFKA_REST_PROXY"
        version = "kafka-rest-proxy:5.0.0"
    }

}
