provider "instaclustr" {
    username = "%s"
    api_key = "%s"
    api_hostname = "%s"
}

resource "instaclustr_cluster" "invalid" {
    cluster_name = "test!@#$"
    node_size = "t3.small"
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
        bundle = "APACHE_CASSANDRA"
        version = "3.11.4"
    }

    bundle {
        bundle = "SPARK"
        version = "apache-spark:2.3.2.ic1"
    }

    bundle {
        bundle = "ZEPPELIN"
        version = "apache-zeppelin:0.8.0-spark-2.3.2"
    }
}
