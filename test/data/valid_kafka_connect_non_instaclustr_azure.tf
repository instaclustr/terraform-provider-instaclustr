provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_cluster" "validKC" {
    cluster_name = "testcluster"
    node_size = "t3.medium-10-gp2"
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
        bundle = "KAFKA_CONNECT"
        version = "2.3.1"
        options = {
            azure_storage_account_name = "%s"
            azure_storage_account_key = "%s"
            azure_storage_container_name = "%s"
            ssl_enabled_protocols = "%s"
            ssl_trustore_password = "%s"
            ssl_protocol = "%s"
            security_protocol = "%s"
            sasl_mechanism = "%s"
            sasl_jaas_config = "%s"
            bootstrap_servers = %s"
            truststore = "%s"
        }
    }
}
